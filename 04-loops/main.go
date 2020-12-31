package main

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int

	fmt.Scanf("%d", &n)

	fmt.Printf("is %d prime? %t", n, IsPrime(n))
}
