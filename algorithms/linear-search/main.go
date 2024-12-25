package linearsearch

func LinearSearch(arr []int, target int) int {

	if len(arr) == 0 {
		return -1
	}

	for i, v := range arr {
		if target == v {
			return i
		}
	}
	return -1
}
