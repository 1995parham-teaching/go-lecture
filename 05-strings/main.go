package main

import "fmt"

// https://en.wikipedia.org/wiki/UTF-8

// g defined as a global constant variable,
// which is added into binary output .rodata section.
// $ strings 05-strings | grep "Global"
const g = "Global string is which defined"

func main() {
	var s1 string = "Hello World"

	s2 := "Parham Alvani"
	s3 := "سلام دنیا"

	fmt.Printf("%d (len(s3)) != 9\n", len(s3))
	fmt.Println([]byte(s3))

	fmt.Println(s1[0])
	fmt.Println(s2)
	// 11011000 (216) means it needs another byte to be a valid UTF-8.
	fmt.Println(s3[0])
	fmt.Println(s3[1])
	fmt.Printf("%c\n", s3[1])

	for i, c := range s3 {
		fmt.Printf("[%d]: %c ", i, c)
	}

	// you can assign a new string
	// into a string variable.
	s1 = "New Hello World"

	// s3[1] = 10
	// cannot assign to s3[1] (value of type byte)

	fmt.Println()

	// comment out the following line and then you cannot find
	// the string into the compiled binary.
	fmt.Println(g)
}
