package main

import "fmt"

func main() {
	arr := []int{11, 11, 11, 22, 33, 33, 44, 55, 55}
	
	newArr := removeDuplicate(arr)
	fmt.Println(newArr)
}

func removeDuplicate(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	k := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[k] = arr[i]
			k++
		}
	}
	return arr[:k]
}