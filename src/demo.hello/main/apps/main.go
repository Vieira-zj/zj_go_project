package main

import (
	"fmt"
	"time"

	"src/demo.hello/apps"
)

func testPool() {
	pool := apps.NewPool(2)
	for i := 0; i < 5; i++ {
		pool.NewTask(func() {
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Println("task done:", time.Now())
		})
	}

	for pool.Size() > 0 {
		fmt.Println("wait, pool size:", pool.Size())
		time.Sleep(time.Second)
	}
	fmt.Println("test pool done.")
}

func main() {
	testPool()
	fmt.Println("apps demo done.")
}
