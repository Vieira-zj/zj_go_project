package examples

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// example 01
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func routineTest() {
	go spinner(100 * time.Millisecond)
	time.Sleep(5 * time.Second)
}

// example 02
func myUpdateRoutine(num int, ch chan string) {
	fmt.Println("start: update goroutine:", num)
	time.Sleep(200 * time.Millisecond)
	ch <- fmt.Sprintf("%s: %d", "test", num) // send
	fmt.Println("end: update goroutine", num)
}

func channelTest() {
	// output 5 goroutine start
	// if no buffered ch, output 3 goroutine end
	// if buffered(1) ch, output 4 goroutine end
	const count = 5
	ch := make(chan string, 1)

	for i := 0; i < count; i++ {
		go myUpdateRoutine(i, ch) // start goroutine
	}

	for i := 0; i < count-2; i++ {
		fmt.Println(<-ch) // receive
	}
	time.Sleep(500 * time.Millisecond)
}

// example 03
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
		go myFetch(url, ch) // start goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// example 03
func selectTest() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("receive at", i)
			fmt.Println(x)
		case ch <- i:
			fmt.Println("send at", i)
		}
	}
}

// MainGoRoutine : main function for goroutine, channel examples.
func MainGoRoutine() {
	// routineTest()
	channelTest()

	// myFetchAllTest()

	// selectTest()

	fmt.Println("\nconcurrent demo.")
}
