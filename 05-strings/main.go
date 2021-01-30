package main

import "fmt"

func main() {
	var s1 string = "Hello World"

	s2 := "Parham Alvani"
	s3 := "سلام دنیا"

	fmt.Printf("%d (len(s3)) != 9\n", len(s3))

	fmt.Println(s1[0])
	fmt.Println(s2)
	fmt.Println(s3[1])
	fmt.Printf("%c\n", s3[1])

	for _, c := range s3 {
		fmt.Printf("%c ", c)
	}

	fmt.Println()
}
