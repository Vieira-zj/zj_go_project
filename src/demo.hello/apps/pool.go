package apps

import (
	"fmt"
	"time"
)

// Pool goroutines pool
type Pool struct {
	work chan func()
	sem  chan struct{}
}

// NewPool creates goroutine instance
func NewPool(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

// Size returns number of goroutine in pool
func (p *Pool) Size() int {
	return len(p.sem)
}

// NewTask adds a task into pool
func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
		fmt.Println("run task on exist worker")
	case p.sem <- struct{}{}:
		fmt.Println("create new worker for task")
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
	}()

	task()
	for {
		select {
		case task = <-p.work:
			fmt.Println("exec task")
			task()
		case <-time.NewTicker(time.Duration(10) * time.Second).C:
			fmt.Println("idle for max time, and exit")
			return
		}
	}
}
