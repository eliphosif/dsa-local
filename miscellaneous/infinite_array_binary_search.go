package miscellaneous

func FindElementInfiniteArray(arr []int, target int) int {
	start := 0
	end := 1

	for target > arr[end] {
		tmp := end + 1
		end = end + 2*(end-start+1)
		start = tmp
	}

	return binarySearch(arr, target, start, end)
}

func binarySearch(arr []int, target int, start int, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
