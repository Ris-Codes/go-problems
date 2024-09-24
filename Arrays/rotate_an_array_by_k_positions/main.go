package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 5 // k-th position
	fmt.Println(rotateArray(arr, k))
}

func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func rotateArray(arr []int, k int) []int {
	k = k % len(arr)

	reverseArray(arr)
	reverseArray(arr[:k])
	reverseArray(arr[k:])
	return arr
}
