package main

import "fmt"

func main() {
	var a [10]int
	var b = [...]int{1, 2, 3, 4}
	c := [4]int{2, 4, 6, 8}
	d := [4]int{1}

	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	for index, value := range b {
		fmt.Printf("b[%d] = %d\n", index, value)
	}

	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// a = c
	// compile error, cannot use c (type [4]int) as type [10]int in assignment
	b = c

	fmt.Println("b:", b)

	b[0] = 10

	fmt.Println("c after chaning b", c)
	fmt.Println("b:", b)

	// n := 10
	// var c [n]int
	// compile error, arrays' length must be constant
}
