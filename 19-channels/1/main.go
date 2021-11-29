package main

import "fmt"

func main() {
	// make(chan int, 10) with 10 rooms
	ch := make(chan int) // with 0 room

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// if you forget closing the channel it causes deadlock
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
