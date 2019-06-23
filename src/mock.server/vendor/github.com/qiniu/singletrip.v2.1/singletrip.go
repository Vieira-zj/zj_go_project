package singletrip

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/qiniu/log.v1"
	"github.com/qiniu/rpc.v1"
	"github.com/qiniu/xlog.v1"
	"qbox.us/rateio"
)

func ExceedReqLimitResp() *http.Response {
	info := "exceed request limit"
	return &http.Response{
		StatusCode:    403,
		Body:          ioutil.NopCloser(strings.NewReader(info)),
		ContentLength: int64(len(info)),
	}
}

type Config struct {
	MaxMemory              int            `json:"max_memory"`
	ReadTimeoutMs          int            `json:"read_timeout_ms"`
	TempDirs               []string       `json:"temp_dirs"`
	DefaultRateLimit       int            `json:"default_rate_limit"`
	RateLimitSizeThreshold int            `json:"rate_limit_size_threshold"`
	RateLimit              map[string]int `json:"rate_limit"`
	DefaultReqLimit        int            `json:"default_req_limit"`
	ReqLimit               map[string]int `json:"req_limit"`
}

type Group struct {
	Config

	mu    sync.Mutex
	calls map[string]*call

	Transport      http.RoundTripper
	CreateTempFile func(dir, prefix string) (*os.File, error)
	tempDirIdx     uint64
	readTimeoutDur time.Duration

	reqLimit  map[string]int
	rateLimit map[string]*rateio.Controller
	limitLock sync.RWMutex
}

type call struct {
	wg    sync.WaitGroup
	resp  *http.Response
	err   error
	nproc int64
	body  buffer
}

func New(conf Config) (*Group, error) {
	if conf.MaxMemory == 0 {
		conf.MaxMemory = 4 << 20 // 4M.
	}
	if conf.ReadTimeoutMs == 0 {
		conf.ReadTimeoutMs = 20000 // 20s.
	}
	if len(conf.TempDirs) == 0 {
		conf.TempDirs = []string{os.TempDir()}
	}
	if conf.DefaultRateLimit <= 0 {
		conf.DefaultRateLimit = 1024 * 1024 * 1024 // 1GB/s
	}
	if conf.DefaultReqLimit <= 0 {
		conf.DefaultReqLimit = 2048
	}
	for _, tempDir := range conf.TempDirs {
		if err := os.MkdirAll(tempDir, 0700); err != nil {
			return nil, err
		}
		cleanTempFiles(tempDir)
	}

	g := &Group{
		calls:          make(map[string]*call),
		Transport:      http.DefaultTransport,
		CreateTempFile: ioutil.TempFile,
		Config:         conf,
		readTimeoutDur: time.Duration(conf.ReadTimeoutMs) * time.Millisecond,
		reqLimit:       make(map[string]int),
		rateLimit:      make(map[string]*rateio.Controller),
	}
	return g, nil
}

func cleanTempFiles(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Warnf("Failed to open dir: %s, err: %v", dir, err)
		return
	}
	defer f.Close()

	fis, err := f.Readdir(-1)
	if err != nil {
		log.Warnf("Failed to readdir: %s, err: %v", dir, err)
		return
	}
	for _, fi := range fis {
		if !fi.IsDir() && strings.HasPrefix(fi.Name(), "singletrip") {
			os.Remove(filepath.Join(dir, fi.Name()))
		}
	}
	return
}

func (g *Group) createTempFile(xl *xlog.Logger, prefix string) (f *os.File, err error) {
	for i := 0; i < len(g.TempDirs); i++ {
		idx := atomic.AddUint64(&g.tempDirIdx, 1) % uint64(len(g.TempDirs))
		dir := g.TempDirs[idx]
		f, err = g.CreateTempFile(dir, prefix)
		if err == nil {
			return
		}
		xl.Warnf("Failed to create temp file, i: %d, err: %v", i, err)
	}
	return
}

func (g *Group) Has(key string) bool {
	g.mu.Lock()
	_, ok := g.calls[key]
	g.mu.Unlock()
	return ok
}

func (g *Group) Do(key string, req *http.Request) (*http.Response, error) {
	r, _, err := g.DoWithClient(key, req, nil)
	return r, err
}

func (g *Group) DoWithClient(key string, req *http.Request, client *rpc.Client) (r *http.Response, isNewReq bool, err error) {
	xl := xlog.NewWithReq(req)
	isNewReq = true
	g.mu.Lock()
	if c, ok := g.calls[key]; ok {
		c.nproc++
		g.mu.Unlock()
		c.wg.Wait()
		isNewReq = false
		if c.err != nil {
			xl.Warnf("Failed to wait round trip, err: %v", c.err)
			return nil, isNewReq, c.err
		}
		xl.Debugf("Wait done in singletrip, key: %s", key)
		return g.newCachedResp(key, c), isNewReq, nil
	}

	c := new(call)
	c.nproc++
	c.wg.Add(1)
	g.calls[key] = c

	g.mu.Unlock()

	limitKey := getHost(req)
	if !g.checkLimit(limitKey) {
		//如果向下执行过程中出错，立即释放限额，否则等到数据拉取完成再释放限额
		g.mu.Lock()
		delete(g.calls, key)
		g.mu.Unlock()
		c.wg.Done()
		return ExceedReqLimitResp(), isNewReq, nil
	}

	if client != nil {
		c.resp, c.err = client.Do(xl, req)
	} else {
		c.resp, c.err = g.Transport.RoundTrip(req)
	}
	if c.err != nil {
		xl.Warnf("Failed to round trip, err: %v", c.err)
		g.mu.Lock()
		delete(g.calls, key)
		g.mu.Unlock()
		c.wg.Done()
		g.releaseLimit(limitKey)
		return nil, isNewReq, c.err
	}

	resp := c.resp

	// 缓存响应体。
	var buf buffer
	cl := resp.ContentLength
	if 0 <= cl && cl <= int64(g.MaxMemory) {
		buf = &byteBuffer{}
	} else {
		f, err := g.createTempFile(xl, "singletrip")
		if err == nil {
			buf = &fileBuffer{f: f}
		} else {
			xl.Warnf("Failed to create cache file, err: %v", err)
			c.err = err
			g.mu.Lock()
			delete(g.calls, key)
			g.mu.Unlock()
			c.wg.Done()
			resp.Body.Close()
			g.releaseLimit(limitKey)
			return nil, isNewReq, err
		}
	}
	c.body = buf

	go func() {
		defer func() {
			resp.Body.Close()
			g.releaseLimit(limitKey)
		}()
		respBody := io.Reader(resp.Body)
		if resp.ContentLength == -1 || resp.ContentLength > int64(g.Config.RateLimitSizeThreshold) {
			c := g.getRateController(limitKey)
			respBody = c.Reader(respBody)
		}
		n, err := io.Copy(buf, respBody)
		if err != nil {
			xl.Errorf("Failed to copy to singletrip body buffer, n: %d/%d, err: %v", n, resp.ContentLength, err)
		}
		buf.FinishWrite(err)
	}()

	c.wg.Done()

	return g.newCachedResp(key, c), isNewReq, nil
}

func (g *Group) newCachedResp(key string, c *call) *http.Response {
	brClose := func() {
		g.mu.Lock()
		c.nproc--
		if c.nproc == 0 {
			delete(g.calls, key)
			c.body.Close()
		}
		g.mu.Unlock()
	}
	br := &bufferReader{
		buf:         c.body,
		readTimeout: g.readTimeoutDur,
		closeFn:     brClose,
	}

	nresp := cloneResp(c.resp)
	nresp.Body = br
	return nresp
}

func (g *Group) getRateController(key string) *rateio.Controller {
	g.limitLock.Lock()
	defer g.limitLock.Unlock()
	c, ok := g.rateLimit[key]
	if !ok {
		if v, ok := g.Config.RateLimit[key]; ok {
			c = rateio.NewController(v)
		} else {
			c = rateio.NewController(g.Config.DefaultRateLimit)
		}
		g.rateLimit[key] = c
	}
	return c
}

func (g *Group) checkLimit(key string) bool {
	g.limitLock.Lock()
	defer g.limitLock.Unlock()
	mirrorLimit := g.Config.DefaultReqLimit
	if limit, ok := g.Config.ReqLimit[key]; ok {
		mirrorLimit = limit
	}
	if g.reqLimit[key] >= mirrorLimit {
		return false
	}
	g.reqLimit[key]++
	return true
}

func (g *Group) releaseLimit(key string) {
	g.limitLock.Lock()
	defer g.limitLock.Unlock()
	now := g.reqLimit[key] - 1
	g.reqLimit[key] = now
	if now == 0 {
		if c, ok := g.rateLimit[key]; ok {
			c.Close()
			delete(g.rateLimit, key)
		}
	}
}

func cloneResp(resp *http.Response) *http.Response {
	nresp := new(http.Response)
	*nresp = *resp
	nresp.Header = cloneHeader(resp.Header)
	return nresp
}

func cloneHeader(src http.Header) http.Header {
	dst := http.Header{}
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
	return dst
}

func getHost(req *http.Request) (domain string) {
	domain = req.Host
	if domain == "" {
		domain = req.URL.Hostname()
	}
	return
}
