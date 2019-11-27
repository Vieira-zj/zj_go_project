package diskusage

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

// DiskUsage includes disk tools.
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

// ------------------------------
// Dir TreeMap
// ------------------------------

// PrintFilesTree prints files tree map for given directory by limited levels.
func (du *DiskUsage) PrintFilesTree(dirPath string, limit int) error {
	return du.printFilesTreeAtCurDir(dirPath, 0, limit)
}

func (du *DiskUsage) printFilesTreeAtCurDir(dirPath string, curLevel, limit int) error {
	if limit != -1 && curLevel >= limit {
		return nil
	}
	if err := du.verifyPath(dirPath); err != nil {
		return err
	}

	fnPrintPrefix := func(level int) {
		for i := level; i > 0; i-- {
			fmt.Print("|\t")
		}
	}

	fInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, info := range fInfos {
		fnPrintPrefix(curLevel)
		if info.IsDir() {
			fmt.Println(info.Name() + "\\")
			du.printFilesTreeAtCurDir(filepath.Join(dirPath, info.Name()), curLevel+1, limit)
		} else {
			fmt.Println(info.Name())
		}
	}
	return nil
}

// ------------------------------
// Dir Disk Space Usage
// ------------------------------

// PrintDirDiskUsage returns disk space usage for given directory.
func (du *DiskUsage) PrintDirDiskUsage(dirPath string) error {
	var (
		filesCount int64
		bytesCount int64
		ch         = make(chan int64)
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
			filesCount++
			bytesCount += fSize
		case <-time.Tick(time.Duration(100) * time.Millisecond):
			du.printSpaceUsage(filesCount, bytesCount)
		}
	}

	log.Printf("total files and disk usage size (%s):\n", dirPath)
	du.printSpaceUsage(filesCount, bytesCount)
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

// ListFiles lists all dirs and files in given directory.
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

func (du *DiskUsage) printSpaceUsage(filesCount, bytesCount int64) {
	gbytes := float64(bytesCount) / 1e9
	if gbytes >= 1.0 {
		log.Printf("%d files\t%.1f GB\n", filesCount, gbytes)
	} else {
		log.Printf("%d files\t%.1f MB\n", filesCount, float64(bytesCount)/1e6)
	}
}

// ------------------------------
// Utils
// ------------------------------

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
