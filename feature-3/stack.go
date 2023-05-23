package feature_3

import "errors"

type Stack struct {
	items []int
	size  int
}

// add comment
func Construct(maxSize int) Stack {
	items := make([]int, 0, maxSize)
	return Stack{items, 0}
}

// add comment
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, errors.New("stack are empty")
	}
	s.size--
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s Stack) Peek() (int, error) {
	if s.size == 0 {
		return -1, errors.New("stack are empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsValid() (bool, error) {
	return len(s.items) != 0, nil
}

func (s *Stack) Zero() {
	s.items = []int{}
}

func (s Stack) Search(val int) int {
	for idx_searching, item := range s.items {
		if val == item {
			return idx_searching
		}
	}
	return -1
}
