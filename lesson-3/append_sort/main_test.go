package main

import (
	"reflect"
	"testing"
)

func TestAppendSort(t *testing.T) {
	sorted := appendSort([]int{5, 3, 6, 8, 1, 2})
	if !(reflect.DeepEqual([]int{1, 2, 3, 5, 6, 8}, sorted)) {
		t.Error("Expected", []int{1, 2, 3, 5, 6, 8}, "Got", sorted)
	}
}

func BenchmarkAppendSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendSort([]int{5, 3, 6, 8, 1, 2})
	}
}
