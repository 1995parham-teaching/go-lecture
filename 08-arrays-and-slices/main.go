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

	// append a new element into our newly created slice.
	s2 = append(s2, 120)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s1 after \"append(s2, 120)\": %v\n", s1)

	// create another slice from middle of s1
	s3 := s1[2:4]
	fmt.Printf("s3: %v, cap(s3): %d, len(s3): %d\n", s3, cap(s3), len(s3))

	a1 := [3]int{1, 2, 3}

	s4 := a1[1:3]
	s4[0] = 20

	fmt.Printf("s4: %v, cap(s4): %d, len(s4): %d\n", s4, cap(s4), len(s4))
	fmt.Printf("a1 after \"s3[0] = 20\": %v\n", a1)

	s5 := s4

	// compile error: slice can only be compared to nil
	// fmt.Printf("s5 ? s4: %v\n", s5 == s4)
	fmt.Printf("s5: %v\n", s5)

	// panic: runtime error: index out of range [3] with length 2
	// fmt.Printf("invalid access to slice \"s5[3]\": %d", s5[3])

	// compile error: invalid array index 3 (out of bounds for 3-element array)
	// fmt.Printf("invalid access to array \"a1[3]\": %d", a1[3])

	fmt.Printf("s5[1:]: %v\n", s5[1:])
}
