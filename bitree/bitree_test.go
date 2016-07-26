package bitree_test

import (
	. "ds/bitree"
	"fmt"
	"testing"
)

func TestBiTree(t *testing.T) {
	bt := New()
	node := NewNode(200, NewNode(1000, nil, nil), NewNode(9999999, nil, nil))
	bt.SetRoot(NewNode(100, node, NewNode(2000, NewNode(99, nil, nil), nil)))
	bt.Traverse(func(v interface{}) interface{} {
		fmt.Println(v)
		return v
	})

	e, _ := node.Elem()
	l, _ := node.LeftChild()
	fmt.Println(e, l)
}
