package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const processRoutines = 2

func main() {
	var ch chan int
	ch = make(chan int)

	var wg sync.WaitGroup

	// create 2 process to process data from the channel
	for id := 0; id < processRoutines; id++ {
		wg.Add(1)
		go func(id int) {
			// each processor prints the data besides its id
			for {
				i, ok := <-ch
				if !ok {
					// if channel is closed, we will close the processor
					fmt.Printf("closed the processosr %d\n", id)
					wg.Done()
					return
				}
				fmt.Printf("process %d in %d\n", i, id)
			}
		}(id)
	}

	shutdown := make(chan int)

	// produce data evey 1 second into the channel
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("we have a new input")
			// if we have any data in shutdonw channel then
			// close ch or write new data to it.
			select {
			case <-shutdown:
				close(ch)
			default:
				ch <- rand.Intn(10)
			}
		}
	}()

	// wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	close(shutdown)

	wg.Wait()
}
