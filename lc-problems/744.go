package problems

func NextGreatestLetter(letters []byte, target byte) byte {

	length := len(letters)
	start, end := 0, length-1

	for start <= end {
		mid := start + (end-start)/2

		if letters[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return letters[start%length]
}
