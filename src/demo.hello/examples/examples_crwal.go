package examples

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// example 01
// Fetch prints the content found at a URL.
func fetch(url string) (respContent []byte) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	return b
}

func testFetch() {
	const url string = "http://gopl.io"
	fmt.Printf("%s\n", fetch(url))
}

// example 02
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

func testFindLinks() {
	const url string = "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println(link)
	}
}

// example 03, ch5-06
// extract makes an HTTP GET request to the specified URL,
// parses the response as HTML, and returns the links in the HTML document.
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
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
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
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

func testExtract() {
	const url = "http://gopl.io"
	links, err := extract(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "extract: %v\n", err)
		os.Exit(1)
	}
	for _, link := range links {
		fmt.Printf("%v\n", link)
	}
}

// example 04, ch8-06
func crawl(url string) []string {
	fmt.Printf("crawl: %s\n", url)
	list, err := extract(url)
	if err != nil {
		log.Printf("extract error: %v\n", err)
	}
	return list
}

// 一个worklist是一个记录了需要处理的元素的队列，每一个元素都是一个需要抓取的URL列表，
// 不过这一次我们用channel代替slice来做这个队列。
// 每一个对crawl的调用都会在他们自己的goroutine中进行并且会把他们抓到的链接发送回worklist.
func testCrawl() {
	urls := []string{"http://gopl.io"}
	worklist := make(chan []string)

	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// example 05
// tokens is a counting semaphore
// used to enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Printf("crawl: %s\n", url)
	tokens <- struct{}{} // acquire a token
	list, err := extract(url)
	<-tokens // release token
	if err != nil {
		log.Printf("extract error: %v\n", err)
	}
	return list
}

func testCrawl2() {
	urls := []string{"http://gopl.io"}
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	n++
	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

// example 06
func testCrawl3() {
	urls := []string{"http://gopl.io"}
	worklist := make(chan []string) // lists of URLs, may have duplicates
	unseenLink := make(chan string) // de-duplicated URLs

	go func() { worklist <- urls }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLink {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLink <- link
			}
		}
	}
}

// MainLinks : the main for crawl links examples
func MainLinks() {
	// testFetch()
	// testFindLinks()
	// testExtract()

	// testCrawl()
	// testCrawl2()
	// testCrawl3()

	fmt.Println("links example done.")
}
