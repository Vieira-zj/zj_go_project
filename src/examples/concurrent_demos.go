package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func concurrentTest() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

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

func main() {
	// concurrentTest()

	// myFetchAllTest()

	fmt.Println("concurrent demo.")
}
