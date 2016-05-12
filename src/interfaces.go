package main

import "fmt"

type Printer interface {
	Print()
}
type Foo struct {
	X, Y int
}
type Bar struct {
	X, Y float64
}

func (f Foo) Print() {
	fmt.Printf("%d %d\n", f.X, f.Y)
}
func (b Bar) Print() {
	fmt.Printf("%g %g\n", b.X, b.Y)
}
