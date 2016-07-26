package bitree

import "errors"

const (
	PRE_ORDER = 1 << iota
	IN_ORDER
	POST_ORDER
)

type TraverseOrder byte

type BiNode struct {
	e interface{}
	l *BiNode
	r *BiNode
}

type BiTree struct {
	root *BiNode
}

func New() *BiTree {
	return &BiTree{}
}

func NewNode(e interface{}, l *BiNode, r *BiNode) *BiNode {
	return &BiNode{e, l, r}
}

func (bt *BiTree) SetRoot(root *BiNode) {
	bt.root = root
}

func (bt *BiTree) Root() *BiNode {
	if bt == nil {
		return nil
	}
	return bt.root
}

func (bt *BiTree) Traverse(handle func(e interface{}) interface{}) {
	bt.root.traverse(handle)
}

func (bn *BiNode) Elem() (interface{}, error) {
	if bn == nil {
		return nil, errors.New("nil BiNode")
	}
	return bn.e, nil
}

func (bn *BiNode) LeftChild() (*BiNode, error) {
	if bn == nil {
		return nil, errors.New("nil BiNode")
	}
	return bn.l, nil
}

func (bn *BiNode) RightChild() (*BiNode, error) {
	if bn == nil {
		return nil, errors.New("nil BiNode")
	}
	return bn.r, nil
}

func (bn *BiNode) traverse(handle func(e interface{}) interface{}) {
	if bn == nil {
		return
	}
	handle(bn.e)
	bn.l.traverse(handle)
	bn.r.traverse(handle)
}
