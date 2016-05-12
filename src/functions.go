package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func main() {
	fmt.Printf("%d\n", Fibonacci(10))
}
