package main

import (
	"log"
	"math/rand"
	"time"
)

func wait() chan int {
	ch := make(chan int)

	go func() {
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Second)

		close(ch)
	}()

	return ch
}

func main() {
	ch1 := wait()
	ch2 := wait()
	ch3 := wait()

	select {
	case <-ch1:
		log.Println(1)
	case <-ch2:
		log.Println(2)
	case <-ch3:
		log.Println(3)
	default:
		// other cases must have results in the moment or the default case will be exectued
		log.Println("failed")
		return
	}

	log.Println("success")
}
