package problems

import "fmt"

func SplitArray(nums []int, k int) int {

	//calculate min and max sum

	minSum, maxSum := 0, 0

	for _, v := range nums {
		if v > minSum {
			minSum = v
		}
		maxSum = maxSum + v
	}

	fmt.Println(minSum, maxSum)

	return 0
}
