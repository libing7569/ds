package list

import "errors"

type List struct {
	head *listNode
	cnt  uint32
}

type listNode struct {
	e    interface{}
	next *listNode
}

func New() *List {
	return &List{}
}

func (l *List) Get(pos uint32) (interface{}, error) {
	cur := l.head
	cnt := pos
	if pos+1 > l.cnt {
		return nil, errors.New("No such element")
	}

	for cnt > 0 {
		cur = cur.next
		cnt--
	}

	return cur.e, nil
}

func (l *List) Ins(pos uint32, e interface{}) error {
	var prev *listNode = nil
	cur := l.head
	cnt := pos
	if pos > l.cnt {
		return errors.New("No such position")
	}

	for cnt > 0 {
		prev = cur
		cur = cur.next
		cnt--
	}

	if prev == nil {
		l.head = &listNode{e, cur}
	} else {
		prev.next = &listNode{e, cur}
	}

	l.cnt++
	return nil
}

func (l *List) Del(pos uint32) (interface{}, error) {
	var prev *listNode = nil
	cur := l.head
	cnt := pos

	if pos+1 > l.cnt {
		return nil, errors.New("No such element")
	}

	for cnt > 0 {
		prev = cur
		cur = cur.next
		cnt--
	}

	if prev == nil {
		l.head = cur.next
	} else {
		prev.next = cur.next
	}

	l.cnt--

	return cur.e, nil
}

func (l *List) Size() uint32 {
	return l.cnt
}

func (l *List) Traverse(handle func(interface{})) {
	cur := l.head

	for cur != nil {
		handle(cur.e)
		cur = cur.next
	}
}

func main() {

}
