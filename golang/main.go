package main

import "fmt"

func main() {

	sorted := insertionSort([]int{
		5, 2, 4, 6, 1, 3,
	}, 6)

	fmt.Print(sorted)
}

func insertionSort(A []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if A[j-1] > A[j] {

				//Moving left to right and right to left in sequential order
				A[j-1], A[j] = A[j], A[j-1]
			}
		}
	}
	return A
}
