package main

import (
	"fmt"

	"github.com/ppdraga/golang_level_1/lesson-2/calc/calcpack"
)

func main() {
	var a, b, op string

	fmt.Print("Input first number: ")
	_, err := fmt.Scanln(&a)
	checkError(err)

	fmt.Print("Input second number: ")
	_, err = fmt.Scanln(&b)
	checkError(err)

	fmt.Print("Input operation: ")
	_, err = fmt.Scanln(&op)
	checkError(err)

	fmt.Println(a, op, b, "=", calcpack.Calc(a, b, op))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
