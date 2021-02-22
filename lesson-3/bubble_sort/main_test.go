package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	sorted := bubbleSort([]int{5, 3, 6, 8, 1, 2})
	if !(reflect.DeepEqual([]int{1, 2, 3, 5, 6, 8}, sorted)) {
		t.Error("Expected", []int{1, 2, 3, 5, 6, 8}, "Got", sorted)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort([]int{5, 3, 6, 8, 1, 2})
	}
}
