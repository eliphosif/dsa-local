package problems

func SearchRange(nums []int, target int) []int {

	firstOccurence, lastOccurence := -1, -1

	firstOccurence = binarySearch(nums, target, true)

	if firstOccurence != -1 {
		lastOccurence = binarySearch(nums, target, false)
	}

	return []int{firstOccurence, lastOccurence}
}

func binarySearch(nums []int, target int, isFirstOccurence bool) int {

	start, end := 0, len(nums)-1
	occurence := -1

	for start <= end {
		mid := start + (end-start)/2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			if isFirstOccurence {
				end = mid - 1
			} else {
				start = mid + 1
			}
			occurence = mid
		}
	}
	return occurence
}
