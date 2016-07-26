package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New(10, func(e1 interface{}, e2 interface{}) bool {
		return e1.(int) < e2.(int)
	})

	h.Push(10)
	h.Push(12)
	h.Push(8)
	h.Push(14)
	h.Push(5)
	h.Traverse(func(e interface{}) {
		fmt.Println(e)
	})
	//fmt.Println("Hello")
	h.Sort()
	h.Traverse(func(e interface{}) {
		fmt.Println(e)
	})

	//for i := 0; i < len(h.b)-1; i++ {
	//if (h.b[i]).(int) < (h.b[i+1]).(int) {
	//t.Error("Heap sort error")
	//}
	//}

	h.Push(10)
	h.Push(11)
	h.Push(8)
	h.Push(14)
	h.Push(4)
	h.Push(5)
	h.Push(2)

	tmp := 0
	for v, err := h.Pop(); err == nil; v, err = h.Pop() {
		if tmp > v.(int) {
			t.Error("Heap pop error")
		}
		tmp = v.(int)
	}
	//fmt.Println(h.Pop())
	//fmt.Println(h.Pop())
	//fmt.Println(h.Pop())
	//fmt.Println(h.Pop())
	//fmt.Println(h.Pop())

}
