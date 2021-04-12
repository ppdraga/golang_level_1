package main

import (
	"fmt"
	"github.com/ppdraga/golang_level_1/incrementor"
)

func main() {

	inc := incrementor.NewIncrementor()

	if err := inc.SetMaximumValue(5); err != nil {
		panic(err)
	}
	fmt.Println(inc)

	for i := 0; i < 12; i++ {
		inc.IncrementNumber()
		fmt.Println(inc.GetNumber())
	}

	err := inc.SetMaximumValue(-5)
	fmt.Println(err.Error())
}
