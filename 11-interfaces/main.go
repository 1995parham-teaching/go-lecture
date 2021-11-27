package main

import "fmt"

type Printer interface {
	Print() string
}

type Person struct {
	Name string
}

func (p Person) Print() string {
	return fmt.Sprintf("Person %s", p.Name)
}

type Student struct {
	Name string
}

func (s Student) Print() string {
	return s.Name
}

func (s Student) Hello() string {
	return "Hello"
}

type Any interface{}

func main() {
	var p Printer
	s := Student{
		Name: "Linus Torvalds",
	}

	s.Hello()

	p = s

	fmt.Println(p.Print())

	// cast from interface to concrete type with panic
	s = p.(Student)

	// cast from interface to concrete type without panic
	_, ok := p.(Person)
	if !ok {
		fmt.Println("p is not a person")
	}
}
