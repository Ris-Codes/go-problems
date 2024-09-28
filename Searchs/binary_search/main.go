package main

import "fmt"

func main() {
	arr := []int{10, 20, 45, 75, 88}
	target := 75

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Println(target, "found at index", result)
	} else {
		fmt.Println("Target element ot found")
	}
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
