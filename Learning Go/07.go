package main

import "fmt"

type Operator func(a, b int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	// funcBuild := func(op Operator) func(int, int) int {
	funcBuild := func(op Operator) Operator {
		return func(a, b int) int {
			return op(a, b)
		}
	}

	add := funcBuild(add)
	sub := funcBuild(sub)

	fmt.Println(add(1, 2) + sub(1, 2))
}
