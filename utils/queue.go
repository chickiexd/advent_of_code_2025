package utils

import (
	"errors"
	"slices"
)

var ErrEmptyQ = errors.New("q is empty")

type Q[T any] struct {
	Data []T
	size int
}

func (q *Q[T]) Set(v []T) {
	q.Data = v
	q.size = len(v)
}

func (q *Q[T]) Put(v T) {
	q.Data = append(q.Data, v)
	q.size++
}

func (q *Q[T]) Pop() (T, error) {
	if q.size > 0 {
		q.size--
		r := q.Data[0]
		q.Data = slices.Delete(q.Data, 0, 1)
		return r, nil
	} else {
		var zero T
		return zero, ErrEmptyQ
	}
}

func (q *Q[T]) Len() int {
	return q.size
}
