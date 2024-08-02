package util

import (
	"container/list"
	"errors"
	"sync"
)

var emptyQueueErr = errors.New("empty queue")

const DefaultQueueSize = 20

// Queue 线程安全的的泛型队列
type Queue[T any] struct {
	list  *list.List
	size  int
	mutex sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: list.New(),
		size: DefaultQueueSize,
	}
}

func (q *Queue[T]) Push(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.list.Len() >= q.size {
		q.list.Remove(q.list.Front())
	}
	q.list.PushBack(item)
}

func (q *Queue[T]) Pop() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.list.Len() == 0 {
		var zero T
		return zero, emptyQueueErr
	}
	element := q.list.Front()
	q.list.Remove(element)
	return element.Value.(T), nil
}

func (q *Queue[T]) Peek() (T, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.list.Len() == 0 {
		var zero T
		return zero, emptyQueueErr
	}
	return q.list.Front().Value.(T), nil
}

func (q *Queue[T]) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.list.Len() == 0
}

func (q *Queue[T]) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.list.Len()
}
