package problems

func Search(nums []int, target int) int {

	pivot := FindPivot(nums)

	if pivot == -1 {
		return binarySearchPivot(nums, 0, len(nums)-1, target)
	}

	if target == nums[pivot] {
		return pivot
	}

	if target >= nums[0] {
		return binarySearchPivot(nums, 0, pivot, target)
	}

	return binarySearchPivot(nums, pivot+1, len(nums)-1, target)
}

func FindPivot(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2

		if mid < end && nums[mid] > nums[mid+1] {
			return mid
		}
		if mid > start && nums[mid] < nums[mid-1] {
			return mid - 1
		}
		if nums[start] >= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func binarySearchPivot(nums []int, start int, end int, target int) int {

	for start <= end {
		mid := start + (end-start)/2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
