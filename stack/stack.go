package stack

import "errors"

type Stack []interface{}

func New(size uint32) Stack {
	return make(Stack, 0, size)
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return s[len(s)-1], nil
}

func (s *Stack) Pop() (e interface{}, err error) {
	if len(*s) == 0 {
		return nil, errors.New("Stack is empty")
	}
	e, err = (*s)[len(*s)-1], nil
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

func (s *Stack) Clear() {
	*s = []interface{}{}
}
