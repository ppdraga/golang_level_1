package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num_str string
	fmt.Print("Input number: ")
	fmt.Scanln(&num_str)
	num, err := strconv.ParseUint(num_str, 10, 16)
	if err != nil {
		fmt.Println("Not a number error", err)
		os.Exit(1)
	}
	if num < 2 {
		fmt.Println("Number must be more then 2", err)
		os.Exit(1)
	}

	var primeNums = []int{2}

	for i := 2; i <= int(num); i++ {
		isPrime := true
		checkLimit := i / 2
		for _, v := range primeNums {
			if i%v == 0 {
				isPrime = false
				break
			}
			if v > checkLimit {
				break
			}
		}
		if isPrime {
			primeNums = append(primeNums, i)
		}
	}

	fmt.Println(primeNums)
}
