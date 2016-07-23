package queue

import "errors"

type Queue struct {
	b []interface{}
}

func New(size uint32) *Queue {
	return &Queue{make([]interface{}, 0, size)}
}

func (q *Queue) En(e interface{}) {
	q.b = append(q.b, e)
}

func (q *Queue) De() (e interface{}, err error) {
	if len(q.b) == 0 {
		return nil, errors.New("Queue is empty")
	}

	e, err = q.b[0], nil
	q.b = q.b[1:]
	return
}

func (q *Queue) Clear() {
	q.b = []interface{}{}
}
