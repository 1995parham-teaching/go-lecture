package main

import "fmt"

// slice s:
// capacity: 5
// array: [ | | | | ]
// size: 2
// s[0] = 0, s[1] = 1
// array: [0|1| | | ]
// s = append(s, 10)
// size: 3
// array: [0|1|10| | ]
// s = append(s, 20)
// capacity: 10
// size: 4
// array: [0|1|10|20| | | | | ..]

func main() {
	s1 := make([]int, 10)

	s1[0] = 10
	s1[1] = 20
	s1[2] = 30

	fmt.Printf("s1: %+v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))

	s1 = append(s1, 10)

	fmt.Printf("s1: %+v, len(s1): %d, cap(s1): %d\n", s1, len(s1), cap(s1))

	for index, value := range s1 {
		fmt.Printf("s1[%d]: %d,", index, value)
	}
	fmt.Println()

	s2 := make([]int, 0, 10)

	s2 = append(s2, 10)

	fmt.Printf("s2: %+v, len(s2): %d, cap(s2): %d\n", s2, len(s2), cap(s2))
}
