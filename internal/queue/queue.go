package queue

import "errors"

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(e T) {
	*q = append(*q, e)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var e T
	if q.Len() == 0 {
		return e, ErrEmptyQueue
	}
	e = (*q)[0]
	*q = (*q)[1:]
	return e, nil
}

func (q *Queue[T]) Len() int {
	return len(*q)
}
