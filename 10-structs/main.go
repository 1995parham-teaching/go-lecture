package main

import "fmt"

type Student struct {
	Name   string
	Family string
	age    int
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Family: %s, age: %d", s.Name, s.Family, s.age)
}

func (s *Student) IncreaseAge(inc int) {
	s.age += inc
}

func main() {
	s := &Student{
		Name:   "Parham",
		Family: "Alvani",
		age:    27,
	}

	fmt.Println(s)

	s.IncreaseAge(10)

	fmt.Println(s)
}
