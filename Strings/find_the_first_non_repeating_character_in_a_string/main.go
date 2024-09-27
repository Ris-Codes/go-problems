package main

import "fmt"

func main() {
	s := "aaabbb"
	if firstNonRepeatingChar(s) != -1 {
		fmt.Println("The first non-repeating character is:", string(firstNonRepeatingChar(s)))
	} else {
		fmt.Println("There are no non-repeating characters")
	}
	
}

func firstNonRepeatingChar(s string) rune {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	for _, ch := range s {
		if freq[ch] == 1 {
			return ch
		}
	}
	return -1
}
