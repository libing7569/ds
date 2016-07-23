package stack

import "errors"

type Stack struct {
	b []interface{}
}

func New(size uint32) *Stack {
	return &Stack{make([]interface{}, 0, 10)}
}

func (s *Stack) Top() interface{} {
	return s.b[len(s.b)-1]
}

func (s *Stack) Pop() (e interface{}, err error) {
	if len(s.b) == 0 {
		return nil, errors.New("Stack is empty")
	}
	e, err = s.b[len(s.b)-1], nil
	s.b = s.b[:len(s.b)-1]
	return
}

func (s *Stack) Push(e interface{}) {
	s.b = append(s.b, e)
}

func (s *Stack) Clear() {
	s.b = []interface{}{}
}
