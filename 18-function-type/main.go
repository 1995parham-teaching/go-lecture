package main

import "fmt"

func logBeforeRun(f func(int) int) func(int) int {
	return func(i int) int {
		fmt.Println("hello, I am log")
		r := f(i)
		fmt.Printf("result: %d\n", r)
		return r
	}
}

func main() {
	f := func(x int) int {
		return x * x
	}

	fmt.Println(f(1))
	fmt.Println(f(2))

	functiosMap := make(map[string]func(int) int)

	functiosMap["square"] = f

	logBeforeRun(f)(10)
}
