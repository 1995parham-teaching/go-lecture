package main

import (
	"fmt"
	"sync"
)

func main() {
	// Use a WaitGroup to wait for the goroutine to finish.
	// Never use `for {}` to wait — it busy-spins a CPU core.
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("goroutine-2")
	})

	fmt.Println("goroutine-1")

	// Block until the goroutine finishes.
	// Without this, main may exit before goroutine-2 ever runs.
	wg.Wait()
}
