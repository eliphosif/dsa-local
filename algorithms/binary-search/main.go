package binarysearch

func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	return binarySearch(arr, start, end, target, arr[end] > arr[start])
}

func binarySearch(arr []int, start int, end int, target int, isAscendingArray bool) int {
	for start <= end {
		mid := start + ((end - start) / 2)

		if arr[mid] == target {
			return mid
		}

		if isAscendingArray {
			if arr[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if arr[mid] > target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
