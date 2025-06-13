// We are going to create a domain specific language for making assertions about arrays.
//
// We want to check if a specific element in an array is equal to some specific value.
//
// The following code should return `true` because the element at index 1 is equal to 33.
//
//   result := forList( []int{22, 33, 44} ).elemAt(1).shouldBe(equalTo(33))
//
// Note: We are not writing a parser. We want to run that code specifically.

// "expected value is 3 at index 1, but got 33".

package main

import "fmt"

func main() {
	result, message := forList([]int{22, 33, 44}).elemAt(1).shouldBe(equalTo(3))
	fmt.Println("Test 1:", result, message)
}

type Condition[T any] interface {
	Check(v Value[T]) (bool, string)
}

type EqualTo[T comparable] struct {
	v T
}

func (e EqualTo[T]) Check(v Value[T]) (bool, string) {
	if e.v != v.v {
		return false, fmt.Sprintf("expected value is %v at index %d, but got %v", e.v, v.index, v.v)
	}

	return true, ""
}

func equalTo[T comparable](v T) EqualTo[T] {
	return EqualTo[T]{v: v}
}

type Value[T any] struct {
	v T

	index int
}

type Values[T any] struct {
	v []T
}

type ValuesMap[K comparable, V any] struct {
	v map[K]V
}

func forList[T any](v []T) Values[T] {
	return Values[T]{
		v: v,
	}
}

func (v Values[T]) elemAt(index int) Value[T] {
	if index >= len(v.v) {
		panic(fmt.Sprintf("your values has only %d elements and you expect something at %d", len(v.v), index+1))
	}

	return Value[T]{
		v:     v.v[index],
		index: index,
	}
}

func (v Value[T]) shouldBe(cnd Condition[T]) (bool, string) {
	return cnd.Check(v)
}
