package problems

import "fmt"

func FindNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		//convert the value into string
		if len(fmt.Sprintf("%v", v))%2 == 0 {
			count++
		}
	}
	return count
}
