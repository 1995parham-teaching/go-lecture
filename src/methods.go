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
	var i integer

	var exm1 Example
	exm1.Val = "Example 1"
	exm1.count = 0

	exm2 := Example{
		Val:   "Example 2",
		count: 10,
	}

	i.log()

	exm1.Log()
	exm1.Log()

	exm2.Log()
	exm2.Log()
}
