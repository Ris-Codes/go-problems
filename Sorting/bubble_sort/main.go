package main

import "fmt"

func main() {
	arr := []int{4, 15, 7, 54, 43, 12}
	fmt.Println("Sorted Array", bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
