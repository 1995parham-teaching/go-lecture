package main

import "fmt"

func main() {

	go func() {
		fmt.Println("goroutine-2")
	}()

	fmt.Println("goroutine-1")

	// block until goroutine execuation
	// everything will be killed when main died
	for {
	}
}
