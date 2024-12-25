package main

import (
	binarysearch "dsa-local/data-structures/binary-search"
	linearsearch "dsa-local/data-structures/linear-search"
	problems "dsa-local/lc-problems"
	"dsa-local/miscellaneous"
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	arr := []int{1, 2, 3, 4, 9, 15, 17, 89, 221}

	fmt.Println(linearsearch.LinearSearch(arr, 10))

	input := [][]int{{1, 2, 3}, {3, 2, 1}}

	fmt.Println(problems.MaximumWealthUsingChan(input))

	target := 5

	fmt.Println(binarysearch.Binarysearch(arr, target))

	letters := []byte{'a', 'b', 'c', 'd', 'y'}
	targetByte := 'c'
	fmt.Println(string(problems.NextGreatestLetter(letters, byte(targetByte))))

	letters = []byte{'x', 'x', 'y', 'y'}
	targetByte = 'z'
	fmt.Println(string(problems.NextGreatestLetter(letters, byte(targetByte))))

	inputNumbers := []int{2, 3, 4, 5, 7, 9, 10}
	targetNumber := 9

	fmt.Println(problems.SearchRange(inputNumbers, targetNumber))

	fmt.Println(miscellaneous.FindElementInfiniteArray(inputNumbers, target))

	mountainArr := []int{2, 3, 4, 5, 7, 9, 10, 5, 3, 1}
	fmt.Println(mountainArr[problems.PeakIndexInMountainArray2(mountainArr)])

	mountainArr2 := problems.MountainArray{
		Array: []int{1, 2, 3, 4, 5, 3, 1},
	}

	fmt.Println(problems.FindInMountainArray(2, &mountainArr2))

	var inputPivot = []int{4, 5, 6, 7, 10, 0, 1, 2}

	fmt.Println(problems.FindPivot(inputPivot))

}
