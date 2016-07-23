package queue

import "errors"

type Queue []interface{}

func New(size uint32) Queue {
	return make(Queue, 0, size)
}

func (q *Queue) En(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) De() (e interface{}, err error) {
	if len(*q) == 0 {
		return nil, errors.New("Queue is empty")
	}

	e, err = (*q)[0], nil
	*q = (*q)[1:]
	return
}

func (q *Queue) Clear() {
	*q = []interface{}{}
}
