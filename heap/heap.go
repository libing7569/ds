package heap

import "errors"

type Heap struct {
	b        []interface{}
	comp     func(interface{}, interface{}) bool
	initSize uint32
}

func New(size uint32, comp func(interface{}, interface{}) bool) *Heap {
	return &Heap{make([]interface{}, 0, size), comp, size}
}

//i start from 1, and the base is start from 0
//calc logic idx is diff from real idx
func (h *Heap) upShift(i int) {
	for ; i/2 > 0; i /= 2 {
		if h.comp(h.b[i-1], h.b[i/2-1]) {
			h.b[i-1], h.b[i/2-1] = h.b[i/2-1], h.b[i-1]
		} else {
			break
		}
	}
}

func (h *Heap) downShift(n int) {
	for i := 1; i*2 <= n; {
		idx := i * 2
		if i*2+1 <= n {
			if h.comp(h.b[i*2], h.b[i*2-1]) {
				idx = i*2 + 1
			}
		}

		if h.comp(h.b[idx-1], h.b[i-1]) {
			h.b[idx-1], h.b[i-1] = h.b[i-1], h.b[idx-1]
			i = idx
		} else {
			break
		}
	}
}

func (h *Heap) buildHeap() {
	for i := 1; i <= len(h.b); i++ {
		h.upShift(i)
	}
}

func (h *Heap) Clear() {
	h.b = make([]interface{}, 0, h.initSize)
}

func (h *Heap) Top() (interface{}, error) {
	if len(h.b) <= 1 {
		return nil, errors.New("Heap is empty")
	}
	return h.b[1], nil
}

func (h *Heap) Pop() (e interface{}, err error) {
	if len(h.b) == 0 {
		return nil, errors.New("Heap is empty")
	}
	e, err = h.b[0], nil
	h.b[0] = h.b[len(h.b)-1]
	h.downShift(len(h.b) - 1)
	h.b = h.b[:len(h.b)-1]
	return
}

func (h *Heap) Push(e int) {
	h.b = append(h.b, e)
	h.upShift(len(h.b))
}

func (h *Heap) Sort() {
	h.buildHeap()
	length := len(h.b)
	for i := 0; i < length-1; i++ {
		h.b[0], h.b[length-i-1] = h.b[length-i-1], h.b[0]
		h.downShift(length - i - 1)
	}
}

func (h *Heap) Traverse(handle func(e interface{})) {
	for _, v := range h.b {
		handle(v)
	}
}
