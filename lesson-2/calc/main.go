package main

import (
	"fmt"
	"os"
	"strconv"
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

	fa, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(a, "is not a number", err)
		os.Exit(1)
	}

	fb, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println(b, "is not a number", err)
		os.Exit(1)
	}

	var res float64
	switch op {
	case "+":
		res = fa + fb
	case "-":
		res = fa - fb
	case "*":
		res = fa * fb
	case "/":
		if fb == 0 {
			fmt.Println("Division by zero error")
			os.Exit(1)
		}
		res = fa / fb
	default:
		fmt.Println(op, "operation is not supported")
		os.Exit(1)
	}

	fmt.Println(fa, op, fb, "=", res)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
