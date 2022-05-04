package main

import "fmt"

// The meaning of 5 factorial is that we need to multiply the numbers from 1 to 5.
// That means, 5! = 5 × 4 × 3 × 2 × 1 = 120.

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	result := factorial(10)
	fmt.Print("Output for Factorial of 5 is ", result)
}
