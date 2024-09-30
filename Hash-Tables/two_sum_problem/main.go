package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 7, 8, 9, 1, 2}
	target := 11
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}
