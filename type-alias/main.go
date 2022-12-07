package main

import "fmt"

type hello = int

type hi int

func main() {
	var h1 hi = 10
	fmt.Println(h1)
	// typecheck: cannot use h1 (variable of type hi) as int value in variable declaration
	// var _ int = h1

	var h2 hello = 10
	fmt.Println(h2)
	var _ int = h2
}
