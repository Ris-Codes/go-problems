package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Current Array: ", arr)

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("Reversed Array: ", arr)
}
