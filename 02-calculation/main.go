package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(1 + 2 + 3)

	var d float64

	d = math.Abs(12.2)
	fmt.Println(d)

	d = math.Abs(-1 * d)
	fmt.Println(d)

	fmt.Printf("%g\n", math.Sin(math.Pi))
	fmt.Printf("%f\n", math.Sin(math.Pi*0.5))
}
