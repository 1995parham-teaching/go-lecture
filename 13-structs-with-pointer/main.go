package main

import "fmt"

type Older interface {
	BeOlder(int)
}

type Person struct {
	Name string
	Age  int
}

func New(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// BeOlder uses reference to Person
func (p *Person) BeOlder(increase int) {
	// there is no ->
	p.Age += increase
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Parham Alvani",
		Age:  27,
	}

	var _ fmt.Stringer = p
	var _ fmt.Stringer = &p

	fmt.Println(p)

	p.BeOlder(1)

	fmt.Println(p)

	var bo Older
	// compile error: cannot use p (type Person) as type Older in assignment:
	// Person does not implement Older (BeOlder method has pointer receiver)
	// bo = p

	bo = &p

	bo.BeOlder(10)

	fmt.Println(p)

	var _ fmt.Stringer = p
	var _ fmt.Stringer = &p

	fmt.Printf("String on Person: %v\n", p)
	fmt.Printf("String on *Person: %v\n", &p)
}
