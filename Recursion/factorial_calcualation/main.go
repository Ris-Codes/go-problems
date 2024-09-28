package main

import "fmt"

func main() {
	n := 10
	fmt.Println("Factorial of", n, "is", factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
