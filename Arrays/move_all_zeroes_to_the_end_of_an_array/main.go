package main

import "fmt"

func main() {
	arr := []int{0, 2, 0, 6, 5, 8, 0, 3}
	fmt.Println("Array after moving all zeroes to end", moveZeroes(arr))
}

func moveZeroes(arr []int) []int {
	nonZeroPosition := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[nonZeroPosition], arr[i] = arr[i], arr[nonZeroPosition]
			nonZeroPosition++
		}
	}
	return arr
}
