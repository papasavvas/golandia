package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	const number = 10
	fmt.Println("Fibonacci of ", number, ":", fibonacci(number))
}
