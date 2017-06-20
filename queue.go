package data

import "fmt"

// A Queue is a collection to which
// elements can be added and from which
// elements can be retrieved. Upon retrieval,
// elements will be removed from the queue.
// In a queue the next retrievable element is
// identified according to the type / implementation
// of the queue, and this element is called the "head"
// of the queue. Upon retrieving an element from a
// queue, it will be removed from the queue.
type Queue interface {
	Insert(e interface{})
	Peek() interface{}
	Pop() interface{}
	Size() int
	IsEmpty() bool
	Clear()
}

type FifoQueue struct {
	queue []interface{}
}

func NewFifoQueue() *FifoQueue {
	return &FifoQueue{queue: make([]interface{}, 0)}
}

func (fq *FifoQueue) Insert(e interface{}) {
	fq.queue = append(fq.queue, e)
}

func (fq *FifoQueue) Peek() interface{} {
	if len(fq.queue) == 0 {
		return nil
	}
	return fq.queue[0]
}

func (fq *FifoQueue) Pop() (e interface{}) {
	if len(fq.queue) == 0 {
		return nil
	}
	e = fq.Peek()
	fq.queue = fq.queue[1:]
	return
}

func (fq *FifoQueue) Size() int {
	return len(fq.queue)
}

func (fq *FifoQueue) IsEmpty() bool {
	return len(fq.queue) == 0
}

func (fq *FifoQueue) Clear() {
	fq.queue = make([]interface{}, 0)
}

func (fq FifoQueue) String() string {
	return fmt.Sprintf("%s", fq.queue)
}

type Stack struct {
	stack []interface{}
	top   int
}

type LifoQueue Stack

func NewStack() *Stack {
	s := &Stack{}
	s.Clear()
	return s
}

func (s *Stack) Insert(e interface{}) {
	s.stack = append(s.stack, e)
	s.top = s.top + 1
}

func (s *Stack) Peek() interface{} {
	if s.top < 0 {
		return nil
	}
	return s.stack[s.top]
}

func (s *Stack) Pop() interface{} {
	if s.top < 0 {
		return nil
	}
	result := s.stack[s.top]
	s.top = s.top - 1
	s.stack = s.stack[:s.top+1]
	return result
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Clear() {
	s.stack = make([]interface{}, 0)
	s.top = -1
}
