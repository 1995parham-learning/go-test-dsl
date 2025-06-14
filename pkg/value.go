package pkg

import "fmt"

type Value[T any] struct {
	v T
}

func (v Value[T]) ShouldBe(cnd Condition[T]) (bool, string) {
	return cnd.Check(v)
}

type Values[T any] struct {
	v []T
}

func (v Values[T]) ElemAt(index int) Value[T] {
	if index >= len(v.v) {
		panic(fmt.Sprintf("your values has only %d elements and you expect something at %d", len(v.v), index+1))
	}

	return Value[T]{
		v: v.v[index],
	}
}

func ForList[T any](v []T) Values[T] {
	return Values[T]{
		v: v,
	}
}

type ValuesMap[K comparable, V any] struct {
	v map[K]V
}

func (v ValuesMap[K, V]) ElemAt(key K) Value[V] {
	if _, ok := v.v[key]; !ok {
		panic(fmt.Sprintf("there is no element with key equal to %v", key))
	}

	return Value[V]{
		v: v.v[key],
	}
}

func ForMap[K comparable, V any](v map[K]V) ValuesMap[K, V] {
	return ValuesMap[K, V]{
		v: v,
	}
}
