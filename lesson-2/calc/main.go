package main

import (
	"fmt"

	"github.com/ppdraga/golang_level_1/lesson-2/calc/calcpack"
)

func main() {
	var a, b, op string

	fmt.Print("Input first number: ")
	fmt.Scanln(&a)

	fmt.Print("Input second number: ")
	fmt.Scanln(&b)

	fmt.Print("Input operation: ")
	fmt.Scanln(&op)

	fmt.Println(a, op, b, "=", calcpack.Calc(a, b, op))
}
