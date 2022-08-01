package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	// m := make(map[int]string)

	m[10] = "Hello"
	m[1373] = "Parham"

	for key, value := range m {
		fmt.Println(key, value)
	}

	var v string

	v = m[10]
	fmt.Printf("m[%d] = %s\n", 10, v)

	v = m[25]
	fmt.Printf("m[%d] = %s\n", 25, v)

	s := make(map[int]bool)

	for i := 0; i < 3; i++ {
		var n int
		fmt.Scanf("%d", &n)

		if s[n] {
			fmt.Printf("%d is already exist\n", n)
		} else {
			s[n] = true
			fmt.Printf("%d saved\n", n)
		}
	}

	opinions := make(map[string]bool)

	opinions["Parham"] = true
	opinions["Sepehr"] = false

	opinion, ok := opinions["Hesam"]
	if ok {
		fmt.Printf("opinions[%s] = %v\n", "Hesam", opinion)
	}
}
