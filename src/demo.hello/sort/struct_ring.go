package sort

import "sync"

/*
DefaultCapacity of an uninitialized Ring buffer.
*/
const DefaultCapacity int = 10

/*
Ring Type implements a Circular Buffer.
The default value of the Ring struct is a valid (empty) Ring buffer with capacity DefaultCapacify.
*/
type Ring struct {
	sync.Mutex
	buf  []interface{}
	head int // the most recent value written
	tail int // the least recent value written
}

/*
SetCapacity sets the maximum size of the ring buffer.
*/
func (r *Ring) SetCapacity(size int) {
	r.Lock()
	defer r.Unlock()

	r.checkInit()
	r.extend(size)
}

/*
Capacity returns the current capacity of the ring buffer.
*/
func (r *Ring) Capacity() int {
	r.Lock()
	defer r.Unlock()

	return r.capacity()
}

/*
ContentSize returns the current number of elements inside the ring buffer.
*/
func (r *Ring) ContentSize() int {
	r.Lock()
	defer r.Unlock()

	if r.head == -1 {
		return 0
	}
	diff := r.head - r.tail
	if diff < 0 {
		diff += r.capacity()
	}
	return diff + 1
}

/*
Enqueue a value into the Ring buffer.
*/
func (r *Ring) Enqueue(i interface{}) {
	r.Lock()
	defer r.Unlock()

	r.checkInit()
	r.set(r.head+1, i) // set by overwrite
	old := r.head
	r.head = r.mod(r.head + 1)
	if old != -1 && r.head == r.tail { // buffer full
		r.tail = r.mod(r.tail + 1)
	}
}

/*
Dequeue a value from the Ring buffer.
Returns nil if the ring buffer is empty.
*/
func (r *Ring) Dequeue() interface{} {
	r.Lock()
	defer r.Unlock()

	r.checkInit()
	if r.head == -1 {
		return nil
	}

	v := r.get(r.tail)
	if r.head == r.tail { // set buffer empty after dequeue
		r.head = -1
		r.tail = 0
	} else {
		r.tail = r.mod(r.tail + 1)
	}
	return v
}

/*
Peek reads the value that Dequeue would have dequeued without actually dequeuing it.
Returns nil if the ring buffer is empty.
*/
func (r *Ring) Peek() interface{} {
	r.Lock()
	defer r.Unlock()

	r.checkInit()
	if r.head == -1 {
		return nil
	}
	return r.get(r.tail)
}

/*
Values returns a slice of all the values in the circular buffer without modifying them at all.
The returned slice can be modified independently of the circular buffer.
However, the values inside the slice are shared between the slice and circular buffer.
*/
func (r *Ring) Values() []interface{} {
	r.Lock()
	defer r.Unlock()

	if r.head == -1 {
		return []interface{}{}
	}

	arr := make([]interface{}, 0, r.capacity())
	for i := 0; i < r.capacity(); i++ {
		idx := r.mod(i + r.tail)
		arr = append(arr, r.get(idx))
		if idx == r.head {
			break
		}
	}
	return arr
}

/**
* Unexported methods beyond this point.
**/

func (r *Ring) capacity() int {
	return len(r.buf)
}

// sets a value at the given unmodified index and returns the modified index of the value
func (r *Ring) set(p int, v interface{}) {
	r.buf[r.mod(p)] = v
}

// gets a value based at a given unmodified index
func (r *Ring) get(p int) interface{} {
	return r.buf[r.mod(p)]
}

// returns the modified index of an unmodified index
func (r *Ring) mod(p int) int {
	return p % len(r.buf)
}

func (r *Ring) checkInit() {
	if r.buf != nil {
		return
	}

	r.buf = make([]interface{}, DefaultCapacity)
	for i := range r.buf {
		r.buf[i] = nil
	}
	r.head, r.tail = -1, 0
}

func (r *Ring) extend(size int) {
	if size == len(r.buf) {
		return
	} else if size < len(r.buf) {
		r.buf = r.buf[:size]
	} else {
		newb := make([]interface{}, size-len(r.buf))
		for i := range newb {
			newb[i] = nil
		}
		r.buf = append(r.buf, newb...)
	}
}
