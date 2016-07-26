package set_test

import (
	. "ds/set"
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Size() != 3 {
		t.Error("Set size error")
	}

	if s.Contains(3) != true {
		t.Error("Set contains error")
	}

	s.Remove(3)

	if s.Contains(3) == true {
		t.Error("Set contains error")
	}
}
