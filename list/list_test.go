package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {

	fmt.Println("Hello")
	l := New()
	l.Ins(0, 10)
	l.Ins(0, 9)
	l.Ins(0, 8)
	l.Ins(0, 7)
	l.Ins(0, 6)
	l.Traverse(func(e interface{}) {
		fmt.Println(e)
	})

	fmt.Println(l.Size())
	fmt.Println(l.Get(3))
	if v, _ := l.Get(3); v != 9 {
		t.Error("Get error")
	}
	l.Del(3)
	l.Traverse(func(e interface{}) {
		fmt.Println(e)
	})

}
