package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	l := largestElement(arr)

	fmt.Println("The largest element in the array: ", l)
}

func largestElement(arr []int) int {
	largest := arr[0]
	for _, n := range arr {
		if n > largest {
			largest = n
		}
	}
	return largest
}
