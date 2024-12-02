package set

import "errors"

var (
	ErrStackOverflow  = errors.New("stack overflow")
	ErrStackUnderflow = errors.New("stack underflow")
)

type Stack struct {
	arr [10]int
	top int
}

func NewStack() *Stack {
	return &Stack{
		arr: [10]int{},
		top: -1,
	}
}

func (s *Stack) Push(k int) error {
	if s.top+1 == len(s.arr) {
		return ErrStackOverflow
	}
	s.top += 1
	s.arr[s.top] = k
	return nil
}

func (s *Stack) Peek() (int, error) {
	if s.top < 0 {
		return 0, ErrStackUnderflow
	}
	return s.arr[s.top], nil
}

func (s *Stack) Pop() (int, error) {
	k, err := s.Peek()
	if err != nil {
		return 0, err
	}
	s.top -= 1
	return k, nil
}
