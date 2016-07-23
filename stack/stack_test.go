package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New(10)
	s1 := []interface{}{1, 2, 3, 4, 5}
	f := func(l []interface{}) {
		s.Clear()
		for _, v := range s1 {
			s.Push(v)
			fmt.Println("push")
		}

		fmt.Println("top")
		if v, _ := s.Top(); v != s1[len(s1)-1] {
			t.Error("Top error")
		}

		fmt.Println("top2")
		i := len(s1) - 1
		for v, ok := s.Pop(); ok == nil; v, ok = s.Pop() {
			fmt.Println(v)
			if s1[i] != v {
				t.Error("Pop error")
			}
			i--
		}
	}
	f(s1)
	f(s1)
}
