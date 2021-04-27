package main

import "fmt"

func main() {
	ch := make(chan int)

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
