package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	Family string `json:"family"`
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

	str := "{\"name\": \"parham\"}"
	var s1 Student

	if err := json.Unmarshal([]byte(str), &s1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)
}
