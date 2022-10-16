package main

import "fmt"

// define new types based on current types

type String string

func (s String) String() string {
	return string(s)
}

type Person struct{}

// define type aliases (just like typedef in c programming language)

type P = Person

type Int = int

func main() {
	var ii Int = 10
	var i int = ii

	fmt.Println(i)

	var ss String = "hello"

	// uncomment the following lines to see:
	// cannot use ss (variable of type String) as string value in assignment
	// var s string
	// s = ss

	fmt.Println(ss)
}
