package main

import "fmt"

// these are global variables which are accessible from functions.
const (
	c1       = 10.1
	c2 int64 = 11
)

var v1 = 10
var v2 uint64 = 10

/*
	var (
		v1 = 10
		v2 uint64 = 10
	)
*/

func main() {
	x := 10 // var x = 10

	var y int

	var z float64 = 10.3

	fmt.Println(x)
	fmt.Printf("%d\n", y)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("%d %d\n", v1, v2)
	fmt.Printf("z = %f\n", z)
	fmt.Printf("z = %.10f\n", z)
	fmt.Printf("z = %g\n", z)
}
