package set

type Empty struct{}
type Set[T comparable] map[T]Empty

func New[T comparable]() *Set[T] {
	return &Set[T]{}
}

func (s *Set[T]) Add(k T) {
	(*s)[k] = Empty{}
}

func (s *Set[T]) Remove(k T) {
	delete(*s, k)
}

func (s *Set[T]) Exists(k T) bool {
	if _, ok := (*s)[k]; ok {
		return true
	}
	return false
}
