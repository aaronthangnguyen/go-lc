package queue

type Queue[T any] []T

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.Empty() {
		return *new(T), false
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e, true
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}
