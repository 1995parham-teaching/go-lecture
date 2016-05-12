package main

import "fmt"

type Example struct {
	Val   string
	count int
}

type integer int

func (i integer) log() {
	fmt.Printf("%d\n", i)
}

func (e *Example) Log() {
	e.count++
	fmt.Printf("%d %s\n", e.count, e.Val)
}

func main() {
}
