package sort

import (
	"fmt"
	"math/rand"
)

// refer: https://github.com/ryszard/goskiplist/blob/master/skiplist

// p is the fraction of nodes with level i pointers that also have level i+1 pointers.
// p equal to 1/4 is a good value from the point of view of speed and space requirements.
// If variability of running times is a concern, 1/2 is a better value for p.
const p = 0.25

const defaultMaxLevel = 32

// ********* Node

// A node is a container for key-value pairs that are stored in a skip list.
type snode struct {
	key, value interface{}
	forward    []*snode
	backward   *snode
}

// next returns the next node in the skip list containing n.
func (n *snode) next() *snode {
	if len(n.forward) == 0 {
		return nil
	}
	return n.forward[0]
}

// previous returns the previous node in the skip list containing n.
func (n *snode) previous() *snode {
	return n.backward
}

// hasNext returns true if n has a next node.
func (n *snode) hasNext() bool {
	return n.next() != nil
}

// hasPrevious returns true if n has a previous node.
func (n *snode) hasPrevious() bool {
	return n.previous() != nil
}

// ********* SkipList

// A SkipList is a map-like data structure that maintains an ordered
// collection of key-value pairs. Insertion, lookup, and deletion are all O(log n) operations.
// A SkipList can efficiently store up to 2^MaxLevel items.
//
// To iterate over a skip list (where s is a *SkipList):
//
//	for i := s.Iterator(); i.Next(); {
//		// do something with i.Key() and i.Value()
//	}
type SkipList struct {
	lessThan func(l, r interface{}) bool
	header   *snode
	footer   *snode
	length   int
	// MaxLevel determines how many items the SkipList can store efficiently (2^MaxLevel).
	//
	// It is safe to increase MaxLevel to accomodate more elements.
	// If you decrease MaxLevel and the skip list already contains nodes on higer levels,
	// the effective MaxLevel will be the greater of the new MaxLevel and the level of the highest node.
	//
	// A SkipList with MaxLevel equal to 0 is equivalent to a standard linked list
	// and will not have any of the nice properties of skip lists (probably not what you want).
	MaxLevel int
}

// Len returns the length of s.
func (s *SkipList) Len() int {
	return s.length
}

// ********* Iterator

// Iterator is an interface that you can use to iterate through the skip list (in its entirety or fragments).
// For an use example, see the documentation of SkipList.
//
// Key and Value return the key and the value of the current node.
type Iterator interface {
	// Next returns true if the iterator contains subsequent elements
	// and advances its state to the next element if that is possible.
	Next() (ok bool)
	// Previous returns true if the iterator contains previous elements
	// and rewinds its state to the previous element if that is possible.
	Previous() (ok bool)
	// Key returns the current key.
	Key() interface{}
	// Value returns the current value.
	Value() interface{}
	// Seek reduces iterative seek costs for searching forward into the Skip List
	// by remarking the range of keys over which it has scanned before.
	// If the requested key occurs prior to the point, the Skip List will start searching as a safeguard.
	// It returns true if the key is within the known range of the list.
	Seek(key interface{}) (ok bool)
	// Close this iterator to reap resources associated with it.
	// While not strictly required, it will provide extra hints for the garbage collector.
	Close()
}

type iter struct {
	current *snode
	key     interface{}
	value   interface{}
	list    *SkipList
}

func (i iter) Key() interface{} {
	return i.key
}

func (i iter) Value() interface{} {
	return i.value
}

func (i *iter) Next() bool {
	if !i.current.hasNext() {
		return false
	}

	i.current = i.current.next()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *iter) Previous() bool {
	if !i.current.hasPrevious() {
		return false
	}

	i.current = i.current.previous()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *iter) Seek(key interface{}) (ok bool) {
	current := i.current
	list := i.list

	// If the existing iterator outside of the known key range,
	// we should set the position back to the beginning of the list.
	if current == nil {
		current = list.header
	}

	// If the target key occurs before the current key,
	// we cannot take advantage of the heretofore spent traversal cost to find it;
	// resetting back to the beginning is the safest choice.
	if current.key != nil && list.lessThan(key, current.key) {
		current = list.header
	}

	if current.backward == nil {
		current = list.header
	} else {
		current = current.backward
	}

	current = list.getPath(current, nil, key)
	if current == nil {
		return
	}

	i.current = current
	i.key = current.key
	i.value = current.value
	return true
}

func (i *iter) Close() {
	i.key = nil
	i.value = nil
	i.current = nil
	i.list = nil
}

// ********* rangeIterator

type rangeIterator struct {
	iter
	upperLimit interface{}
	lowerLimit interface{}
}

func (i *rangeIterator) Next() bool {
	if !i.current.hasNext() {
		return false
	}

	next := i.current.next()
	if !i.list.lessThan(next.key, i.upperLimit) {
		return false
	}

	i.current = i.current.next()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *rangeIterator) Previous() bool {
	if !i.current.hasPrevious() {
		return false
	}

	previous := i.current.previous()
	if i.list.lessThan(previous.key, i.lowerLimit) {
		return false
	}

	i.current = i.current.previous()
	i.key = i.current.key
	i.value = i.current.value
	return true
}

func (i *rangeIterator) Seek(key interface{}) (ok bool) {
	if i.list.lessThan(key, i.lowerLimit) {
		return
	} else if !i.list.lessThan(key, i.upperLimit) {
		return
	}
	return i.iter.Seek(key)
}

func (i *rangeIterator) Close() {
	i.iter.Close()
	i.upperLimit = nil
	i.lowerLimit = nil
}

// ********* SkipList

// Iterator returns an Iterator that will go through all elements s.
func (s *SkipList) Iterator() Iterator {
	return &iter{
		current: s.header,
		list:    s,
	}
}

// Seek returns a bidirectional iterator starting with the first element
// whose key is greater or equal to key; otherwise, a nil iterator is returned.
func (s *SkipList) Seek(key interface{}) Iterator {
	current := s.getPath(s.header, nil, key)
	if current == nil {
		return nil
	}

	return &iter{
		current: current,
		key:     current.key,
		value:   current.value,
		list:    s,
	}
}

// SeekToFirst returns a bidirectional iterator starting from the first element
// in the list if the list is populated; otherwise, a nil iterator is returned.
func (s *SkipList) SeekToFirst() Iterator {
	if s.length == 0 {
		return nil
	}

	current := s.header.next()
	return &iter{
		current: current,
		key:     current.key,
		value:   current.value,
		list:    s,
	}
}

// SeekToLast returns a bidirectional iterator starting from the last element
// in the list if the list is populated; otherwise, a nil iterator is returned.
func (s *SkipList) SeekToLast() Iterator {
	current := s.footer
	if current == nil {
		return nil
	}

	return &iter{
		current: current,
		key:     current.key,
		value:   current.value,
		list:    s,
	}
}

// Range returns an iterator that will go through all the elements of the skip list
// that are greater or equal than from, but less than to.
func (s *SkipList) Range(from, to interface{}) Iterator {
	start := s.getPath(s.header, nil, from)
	return &rangeIterator{
		iter: iter{
			current: &snode{
				forward:  []*snode{start},
				backward: start,
			},
			list: s,
		},
		upperLimit: to,
		lowerLimit: from,
	}
}

func (s *SkipList) level() int {
	return len(s.header.forward) - 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (s *SkipList) effectiveMaxLevel() int {
	return maxInt(s.level(), s.MaxLevel)
}

// Returns a new random level.
func (s SkipList) randomLevel() (n int) {
	for n = 0; n < s.effectiveMaxLevel() && rand.Float64() < p; n++ {
	}
	return
}

// Get returns the value associated with key from s (nil if the key is not present in s).
// The second return value is true when the key is present.
func (s *SkipList) Get(key interface{}) (value interface{}, ok bool) {
	candidate := s.getPath(s.header, nil, key)
	if candidate == nil || candidate.key != key {
		return nil, false
	}
	return candidate.value, true
}

// GetGreaterOrEqual finds the node whose key is greater than or equal to min.
// It returns its value, its actual key, and whether such a node is present in the skip list.
func (s *SkipList) GetGreaterOrEqual(min interface{}) (actualKey, value interface{}, ok bool) {
	candidate := s.getPath(s.header, nil, min)
	if candidate != nil {
		return candidate.key, candidate.value, true
	}
	return nil, nil, false
}

// getPath populates update with nodes that constitute the path to the node that may contain key.
// The candidate node will be returned.
// If update is nil, it will be left alone (the candidate node will still be returned).
// If update is not nil, but it doesn't have enough slots for all the nodes in the path, getPath will panic.
func (s *SkipList) getPath(current *snode, update []*snode, key interface{}) *snode {
	depth := len(current.forward) - 1
	for i := depth; i >= 0; i-- {
		for current.forward[i] != nil && s.lessThan(current.forward[i].key, key) {
			current = current.forward[i]
		}
		if update != nil {
			update[i] = current
		}
	}
	return current.next()
}

// Set sets the value associated with key in s.
func (s *SkipList) Set(key, value interface{}) {
	if key == nil {
		panic("goskiplist: nil keys are not supported")
	}

	// s.level starts from 0, so we need to allocate one.
	update := make([]*snode, s.level()+1, s.effectiveMaxLevel()+1)
	candidate := s.getPath(s.header, update, key)

	if candidate != nil && candidate.key == key {
		candidate.value = value
		return
	}

	newLevel := s.randomLevel()
	// 随机增加level
	if currentLevel := s.level(); newLevel > currentLevel {
		// there are no pointers for the higher levels in update.
		// Header should be there. Also add higher level links to the header.
		for i := currentLevel + 1; i <= newLevel; i++ {
			update = append(update, s.header)
			s.header.forward = append(s.header.forward, nil)
		}
	}

	newNode := &snode{
		forward: make([]*snode, newLevel+1, s.effectiveMaxLevel()+1),
		key:     key,
		value:   value,
	}

	// 给新加入的节点设置前指针
	if previous := update[0]; previous.key != nil {
		newNode.backward = previous
	}

	// 给新加入的节点设置后指针（数组）
	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	s.length++

	if newNode.forward[0] != nil {
		if newNode.forward[0].backward != newNode {
			newNode.forward[0].backward = newNode
		}
	}

	if s.footer == nil || s.lessThan(s.footer.key, key) {
		s.footer = newNode
	}
}

// Delete removes the node with the given key.
// It returns the old value and whether the node was present.
func (s *SkipList) Delete(key interface{}) (value interface{}, ok bool) {
	if key == nil {
		panic("goskiplist: nil keys are not supported")
	}

	update := make([]*snode, s.level()+1, s.effectiveMaxLevel())
	candidate := s.getPath(s.header, update, key)
	if candidate == nil || candidate.key != key {
		return nil, false
	}

	previous := candidate.backward
	if s.footer == candidate {
		s.footer = previous
	}

	// 设置节点的前指针
	next := candidate.next()
	if next != nil {
		next.backward = previous
	}

	// 设置节点levelN的后指针
	for i := 0; i <= s.level() && update[i].forward[i] == candidate; i++ {
		update[i].forward[i] = candidate.forward[i]
	}

	// 删除节点后，levelN链表为空的情况
	for s.level() > 0 && s.header.forward[s.level()] == nil {
		s.header.forward = s.header.forward[:s.level()]
	}

	s.length--

	return candidate.value, true
}

// ********* New

// NewCustomMap returns a new SkipList that will use lessThan as the comparison function.
// lessThan should define a linear order on keys you intend to use with the SkipList.
func NewCustomMap(lessThan func(l, r interface{}) bool) *SkipList {
	return &SkipList{
		lessThan: lessThan,
		header: &snode{
			forward: []*snode{nil},
		},
		MaxLevel: defaultMaxLevel,
	}
}

// ********* Ordered

// Ordered is an interface which can be linearly ordered by the LessThan method,
// whereby this instance is deemed to be less than other.
// Additionally, Ordered instances should behave properly when compared using == and !=.
type Ordered interface {
	LessThan(other Ordered) bool
}

// New returns a new SkipList.
// Its keys must implement the Ordered interface.
func New() *SkipList {
	comparator := func(left, right interface{}) bool {
		return left.(Ordered).LessThan(right.(Ordered))
	}
	return NewCustomMap(comparator)
}

// NewIntMap returns a SkipList that accepts int keys.
func NewIntMap() *SkipList {
	return NewCustomMap(func(l, r interface{}) bool {
		return l.(int) < r.(int)
	})
}

// NewStringMap returns a SkipList that accepts string keys.
func NewStringMap() *SkipList {
	return NewCustomMap(func(l, r interface{}) bool {
		return l.(string) < r.(string)
	})
}

// ********* Set

// Set is an ordered set data structure.
//
// Its elements must implement the Ordered interface.
// It uses a SkipList for storage, and it gives you similar performance guarantees.
//
// To iterate over a set (where s is a *Set):
//
//	for i := s.Iterator(); i.Next(); {
//		// do something with i.Key().
//		// i.Value() will be nil.
//	}
type Set struct {
	skiplist SkipList
}

// NewSet returns a new Set.
func NewSet() *Set {
	comparator := func(left, right interface{}) bool {
		return left.(Ordered).LessThan(right.(Ordered))
	}
	return NewCustomSet(comparator)
}

// NewCustomSet returns a new Set that will use lessThan as the comparison function.
// lessThan should define a linear order on elements you intend to use with the Set.
func NewCustomSet(lessThan func(l, r interface{}) bool) *Set {
	return &Set{skiplist: SkipList{
		lessThan: lessThan,
		header: &snode{
			forward: []*snode{nil},
		},
		MaxLevel: defaultMaxLevel,
	}}
}

// NewIntSet returns a new Set that accepts int elements.
func NewIntSet() *Set {
	return NewCustomSet(func(l, r interface{}) bool {
		return l.(int) < r.(int)
	})
}

// NewStringSet returns a new Set that accepts string elements.
func NewStringSet() *Set {
	return NewCustomSet(func(l, r interface{}) bool {
		return l.(string) < r.(string)
	})
}

// Add adds key to s.
func (s *Set) Add(key interface{}) {
	s.skiplist.Set(key, nil)
}

// Remove tries to remove key from the set. It returns true if key was present.
func (s *Set) Remove(key interface{}) (ok bool) {
	_, ok = s.skiplist.Delete(key)
	return ok
}

// Len returns the length of the set.
func (s *Set) Len() int {
	return s.skiplist.Len()
}

// Contains returns true if key is present in s.
func (s *Set) Contains(key interface{}) bool {
	_, ok := s.skiplist.Get(key)
	return ok
}

// Iterator returns an Iterator that will go through all elements s.
func (s *Set) Iterator() Iterator {
	return s.skiplist.Iterator()
}

// Range returns an iterator that will go through all the elements of the set
// that are greater or equal than from, but less than to.
func (s *Set) Range(from, to interface{}) Iterator {
	return s.skiplist.Range(from, to)
}

// SetMaxLevel sets MaxLevel in the underlying skip list.
func (s *Set) SetMaxLevel(newMaxLevel int) {
	s.skiplist.MaxLevel = newMaxLevel
}

// GetMaxLevel returns MaxLevel fo the underlying skip list.
func (s *Set) GetMaxLevel() int {
	return s.skiplist.MaxLevel
}

// ********* Skiplist Test

func (s *SkipList) printRepr() {
	fmt.Println("header:") // key:nil, value:nil
	for i, link := range s.header.forward {
		if link != nil {
			fmt.Printf("\t%d: -> %v\n", i, link.key)
		} else {
			fmt.Printf("\t%d: -> END\n", i)
		}
	}

	for node := s.header.next(); node != nil; node = node.next() {
		fmt.Printf("%v:%v (level %d)\n", node.key, node.value, len(node.forward))
		for i, link := range node.forward {
			if link != nil {
				fmt.Printf("\t%d: -> %v\n", i, link.key)
			} else {
				fmt.Printf("\t%d: -> END\n", i)
			}
		}
	}
	fmt.Println()
}

// TestSkipList test for skiplist struct.
func TestSkipList() {
	s := NewIntMap()
	s.Set(7, "seven")
	s.Set(1, "one")
	s.Set(0, "zero")
	s.Set(5, "five")
	s.Set(9, "nine")
	s.Set(10, "ten")
	s.Set(3, "three")

	s.printRepr()
	fmt.Println()

	if firstValue, ok := s.Get(0); ok {
		fmt.Println(firstValue)
	}
	fmt.Println()
	// prints:
	//  zero

	s.Delete(7)
	if secondValue, ok := s.Get(7); ok {
		fmt.Println(secondValue)
	} else {
		fmt.Println("key=7 not found!")
	}
	fmt.Println()
	// prints: not found!

	s.Set(9, "niner")
	// Iterate through all the elements, in order.
	unboundIterator := s.Iterator()
	for unboundIterator.Next() {
		fmt.Printf("%d: %s\n", unboundIterator.Key(), unboundIterator.Value())
	}
	fmt.Println()
	// prints:
	//  0: zero
	//  1: one
	//  3: three
	//  5: five
	//  9: niner
	//  10: ten

	for unboundIterator.Previous() {
		fmt.Printf("%d: %s\n", unboundIterator.Key(), unboundIterator.Value())
	}
	fmt.Println()
	//  9: niner
	//  5: five
	//  3: three
	//  1: one
	//  0: zero

	boundIterator := s.Range(3, 10)
	// Iterate only through elements in some range.
	for boundIterator.Next() {
		fmt.Printf("%d: %s\n", boundIterator.Key(), boundIterator.Value())
	}
	fmt.Println()
	// prints:
	//  3: three
	//  5: five
	//  9: niner

	for boundIterator.Previous() {
		fmt.Printf("%d: %s\n", boundIterator.Key(), boundIterator.Value())
	}
	fmt.Println()
	// prints:
	//  5: five
	//  3: three

	var iterator Iterator

	iterator = s.Seek(3)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator = s.Seek(2)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator = s.SeekToFirst()
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  0: zero

	iterator = s.SeekToLast()
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	fmt.Println()
	// prints:
	//  10: ten

	// SkipList can also reduce subsequent forward seeking costs by reusing the same iterator:
	iterator = s.Seek(3)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator.Seek(5)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	fmt.Println()
	// prints:
	//  5: five
}
