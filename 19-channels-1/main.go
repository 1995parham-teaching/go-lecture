package main

import "fmt"

func main() {
	// make(chan int, 10) with 10 rooms (buffered channel)
	ch := make(chan int) // with 0 room (unbuffered channel)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// if you forget to close the channel it causes deadlock.
		// by closing the channel you let Golang to garbage collect it
		// so it if there is anything remaining in the buffered channel it could be
		// removed soon.
		close(ch)
	}()
	// you can read from a channel with
	// i := <-ch

	for i := range ch {
		fmt.Println(i)
	}
}
