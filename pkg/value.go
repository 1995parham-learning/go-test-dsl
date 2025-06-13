package pkg

import "fmt"

type Value[T any] struct {
	v T
}

type Values[T any] struct {
	v []T
}

type ValuesMap[K comparable, V any] struct {
	v map[K]V
}

func ForList[T any](v []T) Values[T] {
	return Values[T]{
		v: v,
	}
}

func (v Values[T]) ElemAt(index int) Value[T] {
	if index >= len(v.v) {
		panic(fmt.Sprintf("your values has only %d elements and you expect something at %d", len(v.v), index+1))
	}

	return Value[T]{
		v: v.v[index],
	}
}

func (v Value[T]) ShouldBe(cnd Condition[T]) (bool, string) {
	return cnd.Check(v)
}
