package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	// Age is a *int so we can distinguish "missing" (nil) from "zero" (0).
	// This is the canonical use-case for the new(expr) form added in Go 1.26.
	Age *int `json:"age,omitempty"`
}

func main() {
	// Go 1.26: new() now accepts an expression for the initial value, so we
	// can build optional pointer fields inline. Previously you needed:
	//
	//     age := 20
	//     s := Student{Name: "Parham", Family: "Alvani", Age: &age}
	s := Student{
		Name:   "Parham",
		Family: "Alvani",
		Age:    new(20),
	}

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// Missing "age": Age stays nil and is omitted from any re-marshal.
	str1 := `{"name": "parham"}`
	var s1 Student

	if err := json.Unmarshal([]byte(str1), &s1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)

	// Explicit "age": 0 — Age is non-nil pointer to 0, distinguishable from missing.
	str2 := `{"name": "parham", "family": "", "age": 0}`
	var s2 Student

	if err := json.Unmarshal([]byte(str2), &s2); err != nil {
		fmt.Println(err)
	}
	if s2.Age != nil {
		fmt.Printf("name=%s family=%q age=%d (present)\n", s2.Name, s2.Family, *s2.Age)
	}
}
