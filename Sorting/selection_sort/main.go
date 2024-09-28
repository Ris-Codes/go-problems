package main

import "fmt"

func main() {
	arr := []int{64, 45, 21, 55, 22}
	fmt.Println("Sorted array", selectionSort(arr))
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
