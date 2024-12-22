package main

import (
	binarysearch "dsa-local/data-structures/binary-search"
	linearsearch "dsa-local/data-structures/linear-search"
	problems "dsa-local/lc-problems"
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	arr := []int{1, 2, 3, 4, 9, 15, 17, 89, 221}

	fmt.Println(linearsearch.LinearSearch(arr, 10))

	input := [][]int{{1, 2, 3}, {3, 2, 1}}

	fmt.Println(problems.MaximumWealthUsingChan(input))

	target := 500

	fmt.Println(binarysearch.Binarysearch(arr, target))

	letters := []byte{'a', 'b', 'c', 'd', 'y'}
	targetByte := 'c'
	fmt.Println(string(problems.NextGreatestLetter(letters, byte(targetByte))))

	letters = []byte{'x', 'x', 'y', 'y'}
	targetByte = 'z'
	fmt.Println(string(problems.NextGreatestLetter(letters, byte(targetByte))))

	inputNumbers := []int{2, 2} //4, 5, 7, 7, 8, 8, 8, 8, 8, 8, 10}
	targetNumber := 2

	fmt.Println(problems.SearchRange(inputNumbers, targetNumber))

}
