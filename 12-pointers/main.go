package main

import "fmt"

// invalid swap because it cannot change its arguments.
func swap(a, b int) {
	a, b = b, a
}

// valid swap because of call by reference.
func swapWithPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 12
	fmt.Printf("before \"swap(%d, %d)\": %d, %d\n", x, y, x, y)
	swap(x, y)
	fmt.Printf("after \"swap(%d, %d)\": %d, %d\n", x, y, x, y)

	fmt.Printf("before \"swapWithPointer(%d, %d)\": %d, %d\n", x, y, x, y)
	swapWithPointer(&x, &y)
	fmt.Printf("after \"swapWithPointer(%d, %d)\": %d, %d\n", x, y, x, y)

	var a int
	a = 10
	fmt.Printf("%d\n", a)

	b := new(int)
	*b = 10
	fmt.Printf("b = %p, *b = %d\n", b, *b)

	var c *int
	fmt.Printf("c: %v\n", c)
}
