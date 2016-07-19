package bitmap

import "sync"

type BitMap struct {
	s []byte
	l sync.RWMutex
}

const (
	BITMAP_INIT_SIZE = 16
	BYTE_SIZE        = 8
)

func New() *BitMap {
	return &BitMap{make([]byte, BITMAP_INIT_SIZE), sync.RWMutex{}}
}

func (b *BitMap) Set(i uint) {
	defer b.l.Unlock()
	b.l.Lock()
	for uint(len(b.s)) < i/BYTE_SIZE+1 {
		b.s = append(b.s, make([]byte, len(b.s))...)
	}
	b.s[i/BYTE_SIZE] |= byte(1 << (i % BYTE_SIZE))
}

func (b *BitMap) Get(i uint) bool {
	if i/BYTE_SIZE+1 > uint(len(b.s)) {
		return false
	}

	defer b.l.RUnlock()
	b.l.RLock()

	return (b.s[i/BYTE_SIZE] & byte(1<<(i%BYTE_SIZE))) != 0
}

func (b *BitMap) Clear() {
	defer b.l.Unlock()
	b.l.Lock()
	b.s = make([]byte, BITMAP_INIT_SIZE)
}
