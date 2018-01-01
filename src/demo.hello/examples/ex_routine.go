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
func spinner1(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func routineTest1() {
	go spinner1(time.Duration(100) * time.Millisecond)
	time.Sleep(time.Duration(5) * time.Second)
}

// example 01-02
func spinner2(ch chan<- string) {
	for i := 0; i < 10; i++ {
		for _, r := range `-\|/` {
			ch <- fmt.Sprintf("\r%c", r)
			time.Sleep(time.Duration(100) * time.Millisecond)
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

// example 01-03, chan type as struct
type typeStrAndInt struct {
	numInt    int
	numString string
}

func myFormatRoutine(num int, ch chan<- typeStrAndInt) {
	var ret typeStrAndInt
	switch num {
	case 1:
		ret = typeStrAndInt{numString: "one", numInt: 1}
	case 2:
		ret = typeStrAndInt{numString: "two", numInt: 2}
	case 3:
		ret = typeStrAndInt{numString: "three", numInt: 3}
	default:
		ret = typeStrAndInt{numString: "nine", numInt: 9}
	}
	time.Sleep(time.Duration(3) * time.Second)
	ch <- ret
}

func routineTest3() {
	const count = 5
	ch := make(chan typeStrAndInt)
	for i := 0; i < count; i++ {
		go myFormatRoutine(i, ch)
	}

	for i := 0; i < count; i++ {
		ret := <-ch
		fmt.Printf("results: %d => %s\n", ret.numInt, ret.numString)
	}
}

// example 02
func myUpdateRoutine(num int, ch chan<- string) {
	fmt.Println("start: update goroutine:", num)
	time.Sleep(time.Duration(200) * time.Millisecond)
	ch <- fmt.Sprintf("test: %d", num) // send
	fmt.Println("end: update goroutine", num)
}

func bufferredChannelTest() {
	// bufferred channel, for goroutine send, output:
	// print 5 goroutine start
	// if no buffered ch, print 3 goroutine end
	// if buffered(2) ch, print 5 goroutine end
	const count = 5
	ch := make(chan string, 2)

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
	time.Sleep(time.Duration(3) * time.Second)
}

// example 04
func channelSelectTest() {
	ch := make(chan int, 1) // buffer 1
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
	tick = time.Tick(time.Duration(100) * time.Millisecond)

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

// MainGoRoutine : main function for goroutine, channel examples.
func MainGoRoutine() {
	// routineTest1()
	// routineTest2()
	// routineTest3()

	// bufferredChannelTest()
	// channelSelectTest()

	// myFetchAllTest()

	// const dir = "/Users/zhengjin"
	// getDirTotalSize1(dir)
	// getDirTotalSize2(dir)

	fmt.Println("\nparallel demo.")
}
