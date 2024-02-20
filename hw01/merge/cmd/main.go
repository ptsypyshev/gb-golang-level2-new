package main

import (
	"fmt"

	"github.com/ptsypyshev/gb-golang-level2-new/hw01/merge/internal/merge"
)

func main() {
	first := []int{1, 3, 5, 9}
	second := []int{1, 2, 4, 5, 8}
	fmt.Println("Merged slice", merge.MergeSortedSlices(first, second))
}