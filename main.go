package main

import (
	"fmt"
	"probsolve/recursion"
)

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(recursion.Merge_sort(unsorted))
}
