package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	myutils "tools.app/utils"
)

// DiskUsage for disk usage tools.
type DiskUsage struct {
	semaphore chan struct{}
	wg        sync.WaitGroup
}

// NewDiskUsage returns a DiskUsage instance.
func NewDiskUsage() *DiskUsage {
	const numParallel = 5
	return &DiskUsage{
		semaphore: make(chan struct{}, numParallel),
	}
}

// PrintDirDiskUsage returns disk space usage for specified directory.
func (du *DiskUsage) PrintDirDiskUsage(dirPath string) error {
	var (
		nfiles, nbytes int64
		ch             = make(chan int64)
	)

	if err := du.verifyPath(dirPath); err != nil {
		return err
	}

	go func() {
		du.wg.Add(1)
		du.walkDir(dirPath, ch)
		du.wg.Wait()
		close(ch)
	}()

LOOP:
	for {
		select {
		case fSize, ok := <-ch:
			if !ok {
				break LOOP
			}
			nfiles++
			nbytes += fSize
		case <-time.Tick(time.Duration(100) * time.Millisecond):
			du.printSpaceUsage(nfiles, nbytes)
		}
	}

	log.Printf("total files and disk usage size (%s):\n", dirPath)
	du.printSpaceUsage(nfiles, nbytes)
	return nil
}

func (du *DiskUsage) walkDir(dirPath string, ch chan<- int64) {
	defer func() {
		if p := recover(); p != nil {
			log.Println("WalkDir routine internal err:", p)
		}
		du.wg.Done()
	}()

	files, err := du.ListFiles(dirPath)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() {
			du.wg.Add(1)
			go du.walkDir(filepath.Join(dirPath, f.Name()), ch)
		} else {
			ch <- f.Size()
		}
	}
}

// ListFiles lists dirs and files in directory.
func (du *DiskUsage) ListFiles(dirPath string) ([]os.FileInfo, error) {
	du.semaphore <- struct{}{}
	defer func() {
		<-du.semaphore
	}()

	if err := du.verifyPath(dirPath); err != nil {
		return nil, err
	}
	return ioutil.ReadDir(dirPath)
}

func (du *DiskUsage) printSpaceUsage(nfiles, nbytes int64) {
	gbytes := float64(nbytes) / 1e9
	if gbytes >= 1.0 {
		log.Printf("%d files\t%.1f GB\n", nfiles, gbytes)
	} else {
		log.Printf("%d files\t%.1f MB\n", nfiles, float64(nbytes)/1e6)
	}
}

func (du *DiskUsage) verifyPath(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("input path argument is empty")
	}

	exist, err := myutils.IsFileExist(path)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("dir/file (%s) is not exist", path)
	}

	return nil
}
