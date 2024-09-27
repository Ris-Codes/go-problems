package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "ohell"
	if isRotation(s1, s2) == true {
		fmt.Println(s1, "is a rotation of", s2)
	} else {
		fmt.Println(s1, "is not a rotation of", s2)
	}
}

func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	combined := s1 + s1
	return strings.Contains(combined, s2)
}
