package main

import "fmt"

func main() {
	nums := []int{-4, 1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	leftPointer, rightPointer := 0, len(nums)-1
	result := make([]int, len(nums))
	leftSquare, rightSquare := 0, 0

	for i := len(nums) - 1; i >= 0; i-- {
		leftSquare = nums[leftPointer] * nums[leftPointer]
		rightSquare = nums[rightPointer] * nums[rightPointer]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			leftPointer++
		} else {
			result[i] = rightSquare
			rightPointer--
		}
	}

	return result
}
