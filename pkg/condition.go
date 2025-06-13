package pkg

import "fmt"

type Condition[T any] interface {
	Check(v Value[T]) (bool, string)
}

type EqualTo[T comparable] struct {
	v T
}

func (e EqualTo[T]) Check(v Value[T]) (bool, string) {
	if e.v != v.v {
		return false, fmt.Sprintf("expected value is %v, but got %v", e.v, v.v)
	}

	return true, ""
}

func EqualToFunc[T comparable](v T) EqualTo[T] {
	return EqualTo[T]{v: v}
}
