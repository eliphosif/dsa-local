package problems

type MountainArray struct {
	Array []int
}

func (ma *MountainArray) get(index int) int {
	return ma.Array[index]
}

func (ma *MountainArray) length() int {
	return len(ma.Array)
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {

	start, end := 0, mountainArr.length()-1

	// First step is to calculate the peak point
	for start < end {
		mid := start + (end-start)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			start = mid + 1
		} else {
			end = mid
		}
	}

	peak := start

	// Search in the ascending part
	leftIndex := binarySearchFindInMountainArray(mountainArr, 0, peak, target, true)
	if leftIndex != -1 {
		return leftIndex
	}

	// Search in the descending part
	return binarySearchFindInMountainArray(mountainArr, peak+1, mountainArr.length()-1, target, false)
}

func binarySearchFindInMountainArray(mountainArr *MountainArray, start int, end int, target int, isAscendingArray bool) int {
	for start <= end {
		mid := start + (end-start)/2
		midValue := mountainArr.get(mid)
		if midValue == target {
			return mid
		}
		if isAscendingArray {
			if midValue < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if midValue < target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}
