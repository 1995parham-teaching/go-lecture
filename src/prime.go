package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	flag := false
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag = true
			break
		}
	}
	if flag {
		fmt.Printf("Not Prime\n")
	} else {
		fmt.Printf("Prime\n")
	}
}
