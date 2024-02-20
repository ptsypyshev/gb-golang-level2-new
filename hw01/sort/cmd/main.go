package main

import (
	"fmt"

	"github.com/ptsypyshev/gb-golang-level2-new/hw01/sort/internal/sort"
)

func main() {
	unsorted := []int{3, 2, 5, 4, 6, 1, 5, 9, 0}
	sort.BubbleSort(unsorted)
	fmt.Println(unsorted)
}