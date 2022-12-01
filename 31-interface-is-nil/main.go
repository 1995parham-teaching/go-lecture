package main

import "fmt"

type Helloer interface {
	SayHello()
}

type Hello struct{}

func (Hello) SayHello() {
}

func main() {
	var h *Hello
	var he Helloer = h

	if he != nil {
		fmt.Println("he is not nil")
	}
}
