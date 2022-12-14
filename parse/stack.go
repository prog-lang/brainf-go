package parse

type stack[T any] struct {
	items []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		items: make([]T, 0),
	}
}

func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) pop() (item T) {
	last := len(s.items) - 1
	item = s.items[last]
	s.items = s.items[:last]
	return
}

func (s *stack[T]) size() int {
	return len(s.items)
}

func (s *stack[T]) empty() bool {
	return s.size() == 0
}
