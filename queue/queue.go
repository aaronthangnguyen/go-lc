package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() T {
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
