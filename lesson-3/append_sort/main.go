package main

import (
	"fmt"
)

func main() {
	unsorted := []int{5, 3, 6, 8, 1, 2}
	sorted := appendSort(unsorted)
	fmt.Println(unsorted)
	fmt.Println(sorted)
}

func appendSort(arr []int) []int {
	arr_result := make([]int, len(arr))
	copy(arr_result, arr)
	arr = arr_result
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			// moving elements to right if they are more then key
			arr[i+1] = arr[i]
			i--
		}
		// append operation
		arr[i+1] = key
	}
	return arr
}
