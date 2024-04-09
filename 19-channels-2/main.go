package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("I am going to write on channel")
		fmt.Println("I am still going to write on channel")
		time.Sleep(time.Second)
		ch <- 10
		fmt.Println("I am done")
		close(ch)
	}()

	time.Sleep(time.Second)
	<-ch
}
