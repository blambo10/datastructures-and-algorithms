package main

import "fmt"

func main() {

	sorted := insertionSort([]int{
		5, 2, 4, 6, 1, 3,
	}, 6)

	fmt.Print(sorted)
}

func IinsertionSort(A []int, n int) []int {
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if A[j-1] > A[j] {
				jBase := j
				jMinus := j - 1

				A[jBase], A[jMinus] = A[jMinus], A[jBase]
			}
		}
	}

	return A
}
