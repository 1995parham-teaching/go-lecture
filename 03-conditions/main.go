package main

import "fmt"

func main() {
	const n = 10

	fmt.Printf("%d\n", Fibonacci(n))

	switch n {
	case 10:
		fmt.Printf("n is equal to 10!\n")
		// uncomment the following line to see what happen
		// fallthrough
	case 11:
		fmt.Printf("in c this statement will be run but here?\n")
		// fallthrough
	default:
		fmt.Println("this should not happen")
	}

	const name = "Parham"

	switch name {
	case "Parham":
		fmt.Println("Yooohoo")
	default:
		fmt.Println("Noooo")
	}
}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
