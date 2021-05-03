package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Printf("%d\n", a)

	b := new(int)
	*b = 10
	fmt.Printf("b = %p, *b = %d\n", b, *b)
}
