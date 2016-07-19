package bitree

import "fmt"
import "testing"

func testBiTree(t *testing.T) {
	bt := New()
	node := &BiNode{1000, &BiNode{9999999, nil, nil}, nil}
	bt.root = &BiNode{100, node, &BiNode{2000, &BiNode{99, nil, nil}, nil}}
	bt.Traverse(func(v interface{}) interface{} {
		fmt.Println(v)
		return v
	})

	e, _ := node.Elem()
	l, _ := node.LeftChild()
	fmt.Println(e, l)
}
