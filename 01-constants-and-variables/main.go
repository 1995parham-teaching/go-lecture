package main

import "fmt"

const (
	c1       = 10.1
	c2 int64 = 11
)

var v1 = 10
var v2 uint64 = 10

func main() {
	x := 10

	var y int

	fmt.Println(x)
	fmt.Printf("%d\n", y)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("%d %d\n", v1, v2)
}
