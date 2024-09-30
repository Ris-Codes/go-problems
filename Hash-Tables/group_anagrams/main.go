package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"aaab", "aaba", "caa", "baaa", "ad", "aac", "da"}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(str []string) [][]string {
	result := make(map[string][]string)

	for _, s := range str {
		sorted := sortString(s)
		result[sorted] = append(result[sorted], s)
	}

	var groupedAnagrams [][]string
	for _, group := range result {
		groupedAnagrams = append(groupedAnagrams, group)
	}
	return groupedAnagrams
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
