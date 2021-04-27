package main

import "fmt"

func main() {
	go func() {
		fmt.Println("chan-2")
	}()

	fmt.Println("chan-1")

	// block until goroutine execuation
	for {
	}
}
