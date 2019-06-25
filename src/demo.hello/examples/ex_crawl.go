package examples

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// example 01
// Fetch prints the content found at a URL.
func testFetchURL() {
	const url string = "http://gopl.io"
	fmt.Printf("\nbytes of url (%s): %d\n", url, len(myFetchURL(url)))
}

func myFetchURL(url string) (respContent []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch url (%s): %v\n", url, err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "io copy: %v\n", err)
		os.Exit(1)
	}
	return b
}

// example 02
func testFetchLinks01() {
	const url string = "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch url (%s): %v\n", url, err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse: %v\n", err)
		os.Exit(1)
	}

	links := visit(nil, doc)
	fmt.Printf("\nall links for url (%s):\n", url)
	for _, link := range links {
		fmt.Println(link)
	}
}

// visit appends to links each link found in html node and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// example 03, ch5-06
func testFetchlLinks02() {
	const url = "http://gopl.io"
	links, err := extract(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "extract: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nall links for url (%s):\n", url)
	for _, link := range links {
		fmt.Printf("%v\n", link)
	}
}

// extract makes an HTTP GET request to the specified URL,
// parses the response as HTML, and returns the links in the HTML document.
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting url(%s): %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html parsing for url(%s): %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return links, nil
}

// forEachNode针对每个结点x, 都会调用pre(x)和post(x). pre和post都是可选的
// 遍历孩子结点之前, pre被调用
// 遍历孩子结点之后, post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// example 04, ch8-06
// 一个worklist是一个记录了需要处理的元素的队列，每一个元素都是一个需要抓取的URL列表，
// 不过这一次我们用channel代替slice来做这个队列。
// 每一个对crawl的调用都会在他们自己的goroutine中进行并且会把他们抓到的链接发送回worklist.
func testCrawl01() {
	urls := []string{"http://gopl.io"}
	chWorklist := make(chan []string)
	go func() { chWorklist <- urls }()

	// crawl web page concurrently (as many as routines)
	seen := make(map[string]bool)
	for links := range chWorklist {
		if isExitCrawlProcess(seen, 100) {
			break
		}
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					chWorklist <- crawl(link)
				}(link)
			}
		}
	}

	fmt.Println("\nall crawl links for urls:", urls)
	for url := range seen {
		fmt.Println(url)
	}
}

func crawl(url string) []string {
	fmt.Printf("crawl links for url: %s\n", url)
	links, err := extract(url)
	if err != nil {
		fmt.Printf("extract error: %v\n", err)
	}
	return links
}

// example 05
func testCrawl02() {
	urls := []string{"http://gopl.io"}
	chWorklist := make(chan []string)
	// tokens is a counting semaphore, used to enforce a limit of x concurrent requests
	var chTokens = make(chan struct{}, 3)
	var n int

	n++
	go func() { chWorklist <- urls }()

	// crawl web page concurrently
	// (as many as routines, and only 5 routines are running)
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		if isExitCrawlProcess(seen, 100) {
			break
		}
		links := <-chWorklist
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					chWorklist <- crawl2(link, chTokens)
				}(link)
			}
		}
	}

	fmt.Println("\nall crawl links for urls:", urls)
	for url := range seen {
		fmt.Println(url)
	}
}

func crawl2(url string, chTokens chan struct{}) []string {
	chTokens <- struct{}{} // acquire a token
	fmt.Printf("crawl: %s\n", url)
	links, err := extract(url)
	if err != nil {
		fmt.Printf("extract error: %v\n", err)
	}
	<-chTokens // release token
	return links
}

// example 06
func testCrawl03() {
	urls := []string{"http://www.baidu.com"}
	// lists of URLs, may have duplicates
	chWorklist := make(chan []string)
	go func() { chWorklist <- urls }()

	// create 5 crawler routines to fetch each unseen link
	chUnseenLink := make(chan string) // de-duplicated URLs
	for i := 0; i < 5; i++ {
		go func() {
			for page := range chUnseenLink {
				// #1: process will be blocked
				// chWorklist <- crawl(page)
				// #2: crawl() is in current routine, and is slow
				links := crawl(page)
				go func() { chWorklist <- links }()
				// #3: put crawl() is in sub routine
				// go func(page string) { chWorklist <- crawl(page) }(page)
			}
		}()
	}

	// the main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
loop:
	for links := range chWorklist {
		if isExitCrawlProcess(seen, 50) {
			break loop
		}
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				chUnseenLink <- link
			}
		}
	}

	fmt.Println("\nall crawl links for urls:", urls)
	for url := range seen {
		fmt.Println(url)
	}
}

func isExitCrawlProcess(urls map[string]bool, limit int) bool {
	if len(urls) > limit {
		fmt.Printf("\nget max urls size: %d\n", limit)
		return true
	}
	return false
}

// MainCrawl : main function for crawl web links examples.
func MainCrawl() {
	// testFetchURL()
	// testFetchlLinks01()
	// testFetchlLinks02()

	// testCrawl01()
	// testCrawl02()
	// testCrawl03()

	fmt.Println("golang crawl example DONE.")
}
