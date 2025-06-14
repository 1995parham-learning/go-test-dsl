# Test Go using DSL

## Introduction

We're building a domain-specific language (DSL) in Go to simplify assertions for data structures like arrays and numbers. The goal is to create a more expressive and readable way to write tests.
For instance, to verify that a specific element in an array holds a particular value, you could write an assertion like this:

```go
result := forList([]int{22, 33, 44}).elemAt(1).shouldBe(equalTo(33))
```

This approach makes the test's intent clear and easy to understand at a glance. In the example [11, 33, 44], the expression would evaluate to true.
