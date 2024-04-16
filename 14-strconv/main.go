package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1230"
	a := 10

	// second argument indicates base and it doesn't
	// accept 0x even when you have 16 as base.
	i, _ := strconv.ParseInt(s, 10, 32)

	// compile error: mismatched types int64 and int
	// a = i + a

	a = int(i) + a

	fmt.Println(i)
	fmt.Println(a)

	// print multiple return values with println
	fmt.Println(strconv.Atoi("123"))

	// Generic ParseInt which is written by me!
	y, _ := ParseInt[int]("562", 16)
	fmt.Println(y)
	fmt.Println(y + a)
}

type Number interface {
	int | int32 | int64 | int16 | int8
}

func ParseInt[T Number](s string, b int) (T, error) {
	v, err := strconv.ParseInt(s, b, 64)
	if err != nil {
		return 0, err
	}

	return T(v), nil
}
