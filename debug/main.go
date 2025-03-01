package main

import "fmt"

func findLocalMinima(arr []int, low, high int) int {
    mid := low + (high-low)/2

    // Check if mid is a local minima
    if (mid == 0 || arr[mid-1] > arr[mid]) && (mid == len(arr)-1 || arr[mid+1] > arr[mid]) {
        return mid
    }

    // If the left neighbor is smaller, then the local minima is in the left half
    if mid > 0 && arr[mid-1] < arr[mid] {
        return findLocalMinima(arr, low, mid-1)
    }

    // Otherwise, the local minima is in the right half
    return findLocalMinima(arr, mid+1, high)
}

func main() {
    arr := []int{11, 40, 13, 60, 9, 8, 7, 6, 5}
    index := findLocalMinima(arr, 0, len(arr)-1)
    fmt.Printf("Local minima is at index %d with value %d\n", index, arr[index])
}