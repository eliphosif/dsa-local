package binarysearch

func Binarysearch(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	//find if the arr is asending or desending
	asen := arr[end] > arr[start]

	for start <= end {
		mid := start + ((end - start) / 2)

		if arr[mid] == target {
			return mid
		}

		if asen {
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
