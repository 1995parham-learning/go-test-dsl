package main

import (
	"fmt"

	"github.com/1995parham-learning/go-test-dsl/pkg"
)

func main() {
	result, message := pkg.ForList([]int{22, 33, 44}).ElemAt(1).ShouldBe(pkg.EqualToFunc(3))
	fmt.Println("Test 1:", result, message)
}
