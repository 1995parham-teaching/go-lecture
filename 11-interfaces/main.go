package main

import "fmt"

type Printer interface {
	Print() string
}

type Increaser interface {
	Increase(int)
}

type Student struct {
	Name string
}

func (s *Student) Increase(inc int) {
	return
}

func (s Student) Print() string {
	return s.Name
}

func (s Student) Hello() string {
	return "Hello"
}

func main() {
	var p Printer
	s := &Student{
		Name: "Hello Student",
	}

	s.Hello()

	p = s

	fmt.Println(p.Print())
}
