package main

import "fmt"

func main() {
	// make(chan int, 10) with 10 rooms (buffered channel)
	ch := make(chan int) // with 0 room (unbuffered channel)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// if you forget closing the channel it causes deadlock
		close(ch)
	}()
	// you can read from a channel with
	// i := <-ch

	for i := range ch {
		fmt.Println(i)
	}
}
