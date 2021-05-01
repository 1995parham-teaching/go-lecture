package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int)

	// create 2 process to process data from the channel
	for id := 0; id < 2; id++ {
		go func(id int) {
			// each processor prints the data besides its id
			for {
				i := <-ch
				fmt.Printf("Process %d in %d\n", i, id)
			}
		}(id)
	}

	// produce data evey 1 second into the channel
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Input")
		ch <- rand.Intn(10)
	}
}
