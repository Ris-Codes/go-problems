package main

import (
	"fmt"
)

func main() {
	input := "aabbbaaa"

	isPalindrome := checkPalindrome(input)

	if isPalindrome {
		fmt.Println(input, "is a Palindrome")
	} else {
		fmt.Println(input, "is not a Palindrome")
	}
}

func checkPalindrome(input string) bool {
	left := 0
	right := len(input)-1
	for left < right {
		if input[left] != input[right] {
			return false
		}
		left++
		right--
	}
	return true
}
