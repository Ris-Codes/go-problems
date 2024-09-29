package main

import "fmt"

func main() {
	arr := []int{21, 89, 32, 59, 56, 86, 78}
	fmt.Println("Sorted array:", quickSort(arr, 0, len(arr)-1))
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i]
	return i + 1
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		i := partition(arr, low, high)
		quickSort(arr, low, i-1)
		quickSort(arr, i+1, high)
	}
	return arr
}
