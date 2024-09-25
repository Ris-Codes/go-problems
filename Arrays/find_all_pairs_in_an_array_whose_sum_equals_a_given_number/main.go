package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 1, 0}

	fmt.Println(findPairs(arr, 5))
}

func findPairs(arr []int, target int) [][2]int {
	complementMap := make(map[int]bool)
	pairs := [][2]int{}
	for _, n := range arr {
		complement := target - n
		if complementMap[complement] {
			pairs = append(pairs, [2]int{complement, n})
		}
		complementMap[n] = true
	}
	return pairs
}
