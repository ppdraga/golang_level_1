// Documentation string for package calcpack
package calcpack

import (
	"fmt"
	"os"
	"strconv"
)

// Documentation string for func Calc in package calcpack
func Calc(a, b, op string) float64 {
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
	return res
}
