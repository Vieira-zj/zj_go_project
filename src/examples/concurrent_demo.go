package examples

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func myFetch(url string, ch chan<- string) {
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
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func myFetchAllTest() {
	urls := []string{"https://golang.org", "http://gopl.io", "https://godoc.org"}
	start := time.Now()

	ch := make(chan string)
	for _, url := range urls {
		go myFetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// MainConcurrent : main function for concurrent demos.
func MainConcurrent() {
	myFetchAllTest()

	fmt.Println("concurrent demo.")
}
