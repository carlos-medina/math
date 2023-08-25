package utils

import (
	"errors"
	"fmt"
)

// Factorial returns the product of all positive integers less than or
// equal to n. If n is negative, Factorial returns an error
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New(fmt.Sprintf("Factorial is defined only for positive number; %v is negative", n))
	}

	return factorial(n), nil
}

func factorial(n int) int {
	if n > 1 {
		f := n * factorial(n-1)
		return f
	}

	return 1
}
