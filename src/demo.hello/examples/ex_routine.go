package examples

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// example 01-01
func testRoutine01() {
	go spinner1(time.Duration(100) * time.Millisecond)
	// routine is killed when "main" is done
	fmt.Println("main sleep for 5 seconds ...")
	time.Sleep(time.Duration(5) * time.Second)
}

func spinner1(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// example 01-02
func testRoutine02() {
	ch := make(chan string)
	go spinner2(ch)

	fmt.Println("main running:")
	// get string from channel until close
	for str := range ch {
		fmt.Printf("%s", str)
	}
}

func spinner2(ch chan<- string) {
	for i := 0; i < 10; i++ {
		for _, r := range `-\|/` {
			ch <- fmt.Sprintf("\r%c", r)
			time.Sleep(time.Duration(100) * time.Millisecond)
		}
	}
	close(ch) // close channel
}

// example 01-03, chan type as struct
type structStrAndInt struct {
	numInt    int
	numString string
}

func testRoutine03() {
	const count = 5
	ch := make(chan structStrAndInt)
	for i := 0; i < count; i++ {
		go myFormatRoutine(i, ch)
	}

	for i := 0; i < count; i++ {
		ret := <-ch
		fmt.Printf("channel output: %d => %s\n", ret.numInt, ret.numString)
	}
}

func myFormatRoutine(num int, ch chan<- structStrAndInt) {
	var ret structStrAndInt
	switch num {
	case 1:
		ret = structStrAndInt{numString: "one", numInt: 1}
	case 2:
		ret = structStrAndInt{numString: "two", numInt: 2}
	case 3:
		ret = structStrAndInt{numString: "three", numInt: 3}
	default:
		ret = structStrAndInt{numString: "nine", numInt: 9}
	}
	time.Sleep(time.Duration(3) * time.Second)
	ch <- ret
}

// example 02
func testBufferredChannel() {
	// bufferred channel, for goroutine send, output:
	// print 5 goroutine "start"
	// if no buffered ch, print 3 goroutine "end"
	// if bufferred(2) ch, print 5 goroutine "end"
	const count = 5
	// ch := make(chan string) // cap=0
	ch := make(chan string, 2) // cap=2
	for i := 0; i < count; i++ {
		go myUpdateRoutine(i, ch) // start goroutine
	}

	// if run_count < 5, ex run_count = 3, 2 goroutine will be pending for send;
	// but when main goroutine finished, these 2 goroutine will be killed.
	// if run_count > 5, ex run_count = 6, all goroutine will be done;
	// but main goroutine will be pending for receive, process pending.
	for i := 0; i < count-2; i++ {
		fmt.Println(<-ch) // receive
	}
	fmt.Println("main sleep for 3 seconds ...")
	time.Sleep(time.Duration(3) * time.Second)
}

func myUpdateRoutine(num int, ch chan<- string) {
	fmt.Println("[myUpdateRoutine] start:", num)
	time.Sleep(time.Duration(200) * time.Millisecond)
	ch <- fmt.Sprintf("number: %d", num) // send
	fmt.Println("[myUpdateRoutine] end:", num)
}

// example 03
func testChannelSelect() {
	ch := make(chan int, 1) // cap=1
	for i := 0; i < 10; i++ {
		fmt.Println("\niterate at: ", i)
		select {
		case x := <-ch:
			fmt.Println("receive idx:", x)
		case ch <- i:
			fmt.Println("send idx:", i)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

// example 04
func myFetchAllTest() {
	urls := []string{"http://baidu.com", "http://gopl.io", "https://godoc.org"}

	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetchURL(url, ch) // start goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("[main] %.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchURL(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("[myFetch]: %.2fs\t%7d\t%s", secs, nbytes, url)
}

// example 05
func listEntries(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on chFileSize.
func walkDir(dir string, chFileSize chan<- int64) {
	for _, entry := range listEntries(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, chFileSize)
		} else {
			chFileSize <- entry.Size()
		}
	}
}

func printSpaceUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func getDirTotalSize01() {
	dir := os.Getenv("HOME")
	chFileSize := make(chan int64)
	go func() {
		walkDir(dir, chFileSize)
		close(chFileSize)
	}()

	var nfiles, nbytes int64
	for size := range chFileSize {
		nfiles++
		nbytes += size
	}
	fmt.Printf("\nspace usage for (%s):\n", dir)
	printSpaceUsage(nfiles, nbytes)
}

func getDirTotalSize02() {
	dir := os.Getenv("HOME")
	chFileSize := make(chan int64)
	go func() {
		walkDir(dir, chFileSize)
		close(chFileSize)
	}()

	// print results periodically
	var tick <-chan time.Time
	tick = time.Tick(time.Duration(100) * time.Millisecond)

	fmt.Printf("\nspace usage for (%s):\n", dir)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-chFileSize:
			if !ok {
				// chFileSize was closed
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printSpaceUsage(nfiles, nbytes)
		}
	}
	printSpaceUsage(nfiles, nbytes)
}

// MainGoRoutine : main function for goroutine, channel examples.
func MainGoRoutine() {
	// testRoutine01()
	// testRoutine02()
	// testRoutine03()

	// testBufferredChannel()
	// testChannelSelect()

	// myFetchAllTest()

	// getDirTotalSize01()
	// getDirTotalSize02()

	fmt.Println("golang routine example DONE.")
}
