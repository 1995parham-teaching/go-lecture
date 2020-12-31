package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	const n = 10

	fmt.Printf("%d\n", Fibonacci(n))

	switch n {
	case 10:
		fmt.Printf("n is equal to 10!\n")
	case 11:
		fmt.Printf("in c this statement will be run but here?\n")
	default:
		fmt.Println("this should not happen")
	}
}
