package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{10, 20, 30, 40, 50, 35, 45}
	fmt.Println("Second Largest element in the Array: ", secondLargest(arr))
}

func secondLargest(arr []int) int {
	larger, second := math.MinInt, math.MaxInt
	for _, n := range arr {
		if n > larger {
			second = larger
			larger = n
		} else if n > second && n != larger {
			second = n
		}
	}
	return second
}
