package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Input")
			ch <- 0
		}
	}()

	for id := 0; id < 2; id++ {
		go func(id int) {
			for {
				i := <-ch
				fmt.Printf("Process %d in %d\n", i, id)
			}
		}(id)
	}

	for {
	}
}
