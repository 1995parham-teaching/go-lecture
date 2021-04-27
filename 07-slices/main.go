package main

import "fmt"

func main() {
	s1 := make([]int, 10)

	s1[0] = 10
	s1[1] = 20
	s1[2] = 30

	fmt.Printf("s1: %+v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))

	for index, value := range s1 {
		fmt.Printf("s1[%d]: %d,", index, value)
	}
	fmt.Println()

	s2 := make([]int, 0, 10)

	fmt.Printf("s2: %+v, len(s2): %d, cap(s2): %d\n", s2, len(s2), cap(s2))
}
