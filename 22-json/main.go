package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	// use the following definition for age to see what happens.
	// Age    *int   `json:"age"`
	Age int `json:"age"`
}

func main() {
	s := Student{
		Name:   "Parham",
		Family: "Alvani",
	}

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	str1 := "{\"name\": \"parham\"}"
	var s1 Student

	if err := json.Unmarshal([]byte(str1), &s1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)

	str2 := "{\"name\": \"parham\", \"family\": \"\", \"age\": 0}"
	var s2 Student

	if err := json.Unmarshal([]byte(str2), &s2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
}
