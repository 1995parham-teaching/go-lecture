package main

import "fmt"

func main() {
	// create a slice with length equals to 10.
	s1 := make([]int, 10)

	s1[0] = 1
	s1[1] = 2

	fmt.Printf("s1: %v\n", s1)

	// create a slice from a slice by slice operator.
	// a newly created slice points to base slice.
	s2 := s1[0:2]
	s2[0] = 10

	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s1 after \"s2[0] = 10\": %v\n", s1)

	a1 := [3]int{1, 2, 3}

	s3 := a1[1:3]
	s3[0] = 20

	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("a1 after \"s3[0] = 20\": %v\n", a1)

	s4 := s3

	// compile error: slice can only be compared to nil
	// fmt.Printf("s4 ? s3: %v\n", s4 == s3)
	fmt.Printf("s4: %v\n", s4)

	// panic: runtime error: index out of range [3] with length 2
	// fmt.Printf("invalid access to slice \"s4[3]\": %d", s4[3])

	// compile error: invalid array index 3 (out of bounds for 3-element array)
	// fmt.Printf("invalid access to slice \"a1[3]\": %d", a1[3])

	fmt.Printf("s4[1:]: %v\n", s4[1:])
}
