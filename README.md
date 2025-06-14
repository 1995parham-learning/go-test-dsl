# Test Go using DSL

## Introduction

We are going to create a domain specific language for making assertions about arrays, numbers, etc. in Go
e.g. we want to check if a specific element in an array is equal to some specific value.
The following code should return `true` because the element at index 1 is equal to 33:

```go
result := forList( []int{22, 33, 44} ).elemAt(1).shouldBe(equalTo(33))
```
