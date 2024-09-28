package main

import "fmt"

func main() {
	arr := []int{10, 20, 45, 75, 88}
	target := 75

	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Println(target, "found at index", result)
	} else {
		fmt.Println("Target element ot found")
	}
}

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
