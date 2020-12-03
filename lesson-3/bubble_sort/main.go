package main

import (
	"fmt"
)

func main() {
	unsorted := []int{5, 3, 6, 8, 1, 2}
	sorted := bubbleSort(unsorted)
	fmt.Println(unsorted)
	fmt.Println(sorted)
}

func bubbleSort(arr []int) []int {
	arr_result := make([]int, len(arr))
	copy(arr_result, arr)
	arr = arr_result
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
