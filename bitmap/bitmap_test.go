package bitmap_test

import (
	. "ds/bitmap"
	"math/rand"
	"testing"
)

func TestBitMap(t *testing.T) {

	ts := []uint{}
	bm := New()
	for i := 0; i < 100; i++ {
		tmp := uint(rand.Uint32())
		bm.Set(tmp)
		ts = append(ts, tmp)
	}

	for _, v := range ts {
		if bm.Get(v) == false {
			t.Error("Check Error", v)
		}
	}

	bm.Clear()

	for _, v := range ts {
		if bm.Get(v) == true {
			t.Error("Check Error", v)
		}
	}

}
