package queue

import "testing"

func TestQueue(t *testing.T) {

	q := New(10)
	s := []interface{}{1, 2, 3, 4, 5}
	f := func(s []interface{}) {
		q.Clear()
		for _, v := range s {
			q.En(v)
		}

		i := 0
		for v, ok := q.De(); ok == nil; v, ok = q.De() {
			if v != s[i] {
				t.Error("De error")
			}
			i++
		}
	}

	f(s)
	f(s)

}
