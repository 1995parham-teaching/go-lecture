package main

import "fmt"

type t1 = int

type t2 int

func main() {
	var h1 t2 = 10
	fmt.Println(h1)
	// typecheck: cannot use h1 (variable of type hi) as int value in variable declaration
	// var _ int = h1

	var h2 t1 = 10
	fmt.Println(h2)
	var _ int = h2
}
