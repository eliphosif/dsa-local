package problems

import "fmt"

func PeakIndexInMountainArray2(arr []int) int {

	start, end := 0, len(arr)-1

	for start < end {

		mid := start + (end-start)/2
		arrMidVal := arr[mid]
		fmt.Println(arrMidVal)
		if arr[mid] > arr[mid+1] {
			end = mid
		} else if arr[mid] < arr[mid+1] {
			start = mid + 1
		}

	}
	return start
}
