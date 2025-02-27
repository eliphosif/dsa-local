package problems

import "fmt"

func SortedSquares(nums []int) []int {

	leftPointer, rightPointer := 0, len(nums)-1

	result := []int{}
	leftSquare, rightSquare := 0, 0

	for leftPointer <= rightPointer {

		leftSquare = nums[leftPointer] * nums[leftPointer]
		leftSquare = nums[rightPointer] * nums[rightPointer]

		fmt.Println(leftSquare, rightSquare)
		fmt.Println(leftPointer, rightPointer)
		if leftSquare > rightSquare {
			result = append(result, leftSquare)
			leftPointer++
		} else {
			result = append(result, rightSquare)
			rightPointer--
		}
	}

	return result
}
