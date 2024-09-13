package queue

import (
	"math/rand"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Len() != 0 {
		t.Fatal("bruh")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	N := 1000
	q := NewQueue[int]()
	data := rand.Perm(N)
	for _, num := range data {
		q.Enqueue(num)
	}
	i := 0
	for q.Len() > 0 {
		if num, err := q.Dequeue(); err != nil {
			t.Fatalf("bruh %v", err)
		} else if num != data[i] {
			t.Fatalf("expected %d, got %d", data[i], num)
		} else {
			i++
		}
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := NewQueue[map[string]int]()
	if _, err := q.Dequeue(); err != ErrEmptyQueue {
		t.Fatalf("expected ErrEmptyQueue, got %v", err)
	}
}
