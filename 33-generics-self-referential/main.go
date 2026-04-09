package main

import "fmt"

// Self-referential generic types are a Go 1.26 language change: a generic
// type may now refer to itself in its own type parameter list.
//
// Adder is the classic motivating example — it constrains A to be a type
// that knows how to Add another A and return an A. Before Go 1.26 this
// constraint could not be expressed directly.
type Adder[A Adder[A]] interface {
	Add(A) A
}

// Vec2 is a concrete type that satisfies Adder[Vec2].
type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{X: v.X + o.X, Y: v.Y + o.Y}
}

// Money also satisfies Adder[Money], showing the constraint is reusable.
type Money int64

func (m Money) Add(o Money) Money {
	return m + o
}

// sum works for any Adder, leaning entirely on the self-referential bound.
func sum[A Adder[A]](xs ...A) A {
	var acc A
	for _, x := range xs {
		acc = acc.Add(x)
	}
	return acc
}

func main() {
	fmt.Println(sum(Vec2{1, 2}, Vec2{3, 4}, Vec2{5, 6}))
	fmt.Println(sum(Money(100), Money(250), Money(75)))
}
