package main

import (
	linearsearch "dsa-local/ds/linear-search"
	problems "dsa-local/lc-problems"
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	arr := []int{1, 2, 3, 4}
	fmt.Println(linearsearch.LinearSearch(arr, 10))

	input := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(input)

	fmt.Println(problems.MaximumWealthUsingChan(input))

}
