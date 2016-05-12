package main

import "fmt"

type empty interface{}

type example struct {
	A int
}

func main() {
	one := 1
	var i empty = one
	j := example{A: 10}
	var k empty = j
	fmt.Println(k.(example).A)
}
