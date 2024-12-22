package problems

func PeakIndexInMountainArray(arr []int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] > arr[mid+1] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}
