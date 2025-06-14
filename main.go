package main

import (
	"fmt"

	"github.com/1995parham-learning/go-test-dsl/pkg"
)

func main() {
	result, message := pkg.ForList([]int{22, 33, 44}).ElemAt(1).ShouldBe(pkg.EqualToFunc(3))
	if result {
		fmt.Println("Test 1: Success")
	} else {
		fmt.Printf("Test 1: Failed %s\n", message)
	}

	result, message = pkg.ForList([]int{22, 33, 44}).ElemAt(1).ShouldBe(pkg.EqualToFunc(33))
	if result {
		fmt.Println("Test 1: Success")
	} else {
		fmt.Printf("Test 1: Failed %s\n", message)
	}

	result, message = pkg.ForMap(map[string]int{"k1": 22, "k2": 33, "k3": 44}).ElemAt("k2").ShouldBe(pkg.EqualToFunc(33))
	if result {
		fmt.Println("Test 1: Success")
	} else {
		fmt.Printf("Test 1: Failed %s\n", message)
	}

	result, message = pkg.ForMap(map[string]int{"k1": 22, "k2": 33, "k3": 44}).ElemAt("k2").ShouldBe(pkg.EqualToFunc(3))
	if result {
		fmt.Println("Test 1: Success")
	} else {
		fmt.Printf("Test 1: Failed %s\n", message)
	}
}
