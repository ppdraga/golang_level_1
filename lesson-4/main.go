package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n_str string
	fmt.Print("Input number: ")
	_, err := fmt.Scanln(&n_str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	num, err := strconv.ParseInt(n_str, 10, 32)
	if err != nil {
		fmt.Println(n_str, "is not a number", err)
		os.Exit(1)
	}

	result := fibonachi(1, 1, num)
	fmt.Println(result)

	cached_fibonachi := cached(fibonachi)
	cached_fibonachi(1, 1, 7)
	cached_fibonachi(1, 1, 7)
	cached_fibonachi(1, 1, 7)
	cached_fibonachi(1, 1, 8)
	cached_fibonachi(1, 1, 8)
}

func fibonachi(a, b, n int64) int64 {
	if n == 1 {
		return a
	} else if n == 2 {
		return b
	} else {
		return fibonachi(a, b, n-2) + fibonachi(a, b, n-1)
	}
}

func cached(f func(a, b, n int64) int64) func(a, b, n int64) int64 {
	type args struct {
		a int64
		b int64
		n int64
	}
	cache := make(map[args]int64)
	return func(a, b, n int64) int64 {
		f_args := args{a, b, n}
		if val, ok := cache[f_args]; ok {
			fmt.Println("from cache")
			return val
		} else {
			result := f(a, b, n)
			fmt.Println("to cache")
			cache[f_args] = result
			return result
		}
	}
}
