package problems

func MoveZeroes(nums []int) {

	leftPointer, rightPointer := 0, 0

	for rightPointer < len(nums) {

		if nums[rightPointer] != 0 {
			nums[leftPointer], nums[rightPointer] = nums[rightPointer], nums[leftPointer]
			leftPointer++
		}

		rightPointer++
	}

}
