package main

import "fmt"

func main() {
	s := "asus"
	freq  := charFreq(s)

	for key, value := range freq {
		fmt.Println( string(key), ":",value)
	}
}

func charFreq(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	return freq
}
