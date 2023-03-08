package queue_test

import (
	"github.com/aaronthangnguyen/go-lc"
	"testing"
)

func TestQueueEmpty(t *testing.T) {
	q := queue.New[int]()

	if !q.Empty() {
		t.Fatal("Queue should be empty")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Empty() {
		t.Fatal("Queue should not be empty")
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if !q.Empty() {
		t.Fatal("Queue should be empty")
	}
}

func TestQueueFIFO(t *testing.T) {
	q := queue.New[int]()

	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 1000; i++ {
		e, _ := q.Dequeue()
		if e != i {
			t.Fatalf("Dequeued element should be %d, got %d", i, e)
		}
	}
}

func TestDequeue(t *testing.T) {
	q := queue.New[int]()

	_, ok := q.Dequeue()
	if ok {
		t.Fatal("Dequeue should not return `ok`")
	}

	q.Enqueue(4)
	e, ok := q.Dequeue()
	if !ok {
		t.Fatal("Dequeue should return `ok`")
	}
	if e != 4 {
		t.Fatalf("Dequeued element should be 4, got %d", e)
	}
}
