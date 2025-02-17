package helpers

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	r := s.items[l]
	s.items = s.items[:l]
	return r
}
