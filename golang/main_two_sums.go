package main

import "fmt"

func main() {
	//fmt.Println("hello world")
	s := []int{3, 2, 4}
	t := 6
	result := twoSum(s, t)

	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	searchIndex := 0
	var result []int

	result = findNumbers(nums, target, searchIndex)

	return result
}

func findNumbers(nums []int, target int, searchIndex int) []int {
	var result []int
	for i, num := range nums {

		if searchIndex < len(nums) {

			if searchIndex != i {
				currentCalculation := nums[searchIndex] + num
				if currentCalculation == target {
					result = append(result, searchIndex, i)
					return result
				}
			}

		}
	}

	searchIndex += 1
	result = findNumbers(nums, target, searchIndex)

	return result
}
