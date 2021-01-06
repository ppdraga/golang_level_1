package calcpack

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {

	type test struct {
		data   []string
		answer float64
	}

	tests := []test{
		test{[]string{"5", "6", "*"}, 30},
		test{[]string{"55", "5", "/"}, 11},
		test{[]string{"15", "6", "-"}, 9},
		test{[]string{"5", "16", "+"}, 21},
	}

	for _, v := range tests {
		x := Calc(v.data[0], v.data[1], v.data[2])
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleCalc() {
	fmt.Println(Calc("5", "6", "*"))
	// Output:
	// 30
}
