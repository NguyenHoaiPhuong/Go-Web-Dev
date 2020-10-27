package main

import (
	"container/list"
	"fmt"
)

// Queue struct
type Queue struct {
	queue *list.List
}

// New func
func NewQueue() *Queue {
	q := new(Queue)
	q.queue = list.New()
	return q
}

// Enqueue : add elem at last
func (q *Queue) Enqueue(v interface{}) {
	q.queue.PushBack(v)
}

// Dequeue : remove last elem
func (q *Queue) Dequeue() {
	elem := q.queue.Back()
	q.queue.Remove(elem)
}

// RemoveAll : remove all elems
func (q *Queue) RemoveAll() {
	for q.queue.Len() > 0 {
		q.Dequeue()
	}
}

// Display : print all elems
func (q *Queue) Display() {
	elem := q.queue.Front()
	for elem != nil {
		fmt.Println(elem.Value)
		elem = elem.Next()
	}
}

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Dequeue()
	q.Display()
}
