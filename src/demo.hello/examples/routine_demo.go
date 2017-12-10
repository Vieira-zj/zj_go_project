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

// example 01
func spinner1(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func routineTest1() {
	go spinner1(100 * time.Millisecond)
	time.Sleep(5 * time.Second)
}

func spinner2(ch chan<- string) {
	for i := 0; i < 10; i++ {
		for _, r := range `-\|/` {
			ch <- fmt.Sprintf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
	close(ch) // close channel
}

func routineTest2() {
	ch := make(chan string)
	go spinner2(ch)

	fmt.Println("running:")
	// get values for channel until close(ch)
	for str := range ch {
		fmt.Printf("%s", str)
	}
}

// example 02
func myUpdateRoutine(num int, ch chan<- string) {
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

// example 04
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

// example 05
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func getDirTotalSize1(dir string) {
	fileSizes := make(chan int64)
	go func() {
		walkDir(dir, fileSizes)
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func getDirTotalSize2(dir string) {
	fileSizes := make(chan int64)
	go func() {
		walkDir(dir, fileSizes)
		close(fileSizes)
	}()

	// print the results periodically
	var tick <-chan time.Time
	tick = time.Tick(100 * time.Millisecond)

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes: // use ok instead of range
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes) // final totals
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// TODO: ch8-09, abort

// MainGoRoutine : main function for goroutine, channel examples.
func MainGoRoutine() {
	// routineTest1()
	// routineTest2()

	// channelTest()
	// selectTest()

	// myFetchAllTest()

	// const dir = "/Users/zhengjin"
	// getDirTotalSize1(dir)
	// getDirTotalSize2(dir)

	fmt.Println("\nconcurrent demo.")
}
