/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 27-03-2021
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import "fmt"

func main() {
	var a [10]int
	var b = [...]int{1, 2, 3, 4}

	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	for index, value := range b {
		fmt.Printf("b[%d] = %d\n", index, value)
	}

	// n := 10
	// var c [n]int
	// compile error, arrays' length must be constant
}
