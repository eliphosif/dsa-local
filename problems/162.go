package problems

func FindPeakElement(nums []int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		} else if nums[mid] > nums[mid+1] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}
