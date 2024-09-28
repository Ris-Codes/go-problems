package main

import "fmt"

func main() {
	arr := []int{34, 56, 45, 53, 36}
	fmt.Println("Sorted array:", insertionSort(arr))
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
