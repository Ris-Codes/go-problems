package main

import "fmt"

func main() {
	s1 := "listen"
	s2 := "silent"
	if areAnagrams(s1, s2) == true {
		fmt.Println(s1, "&", s2, "are anagrams")
	} else {
		fmt.Println(s1, "&", s2, "are not anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	freq := make(map[rune]int)
	for _, ch := range s1 {
		freq[ch]++
	}
	for _, ch := range s2 {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}
	return true
}
