package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findMissingNum(arr, 3))
}

func findMissingNum(arr []int, n int) int {
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}
	return expectedSum - actualSum
}
