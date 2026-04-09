package main

import (
	"fmt"
	"sync"
)

// LazyInit demonstrates sync.OnceValue (Go 1.21+),
// which is the modern, idiomatic replacement for the older
// sync.Once + sentinel-field pattern.
type LazyInit struct {
	value func() int
}

func NewLazyInit() *LazyInit {
	return &LazyInit{
		value: sync.OnceValue(func() int {
			return 1820
		}),
	}
}

func (s *LazyInit) Value() int {
	return s.value()
}

func main() {
	l := NewLazyInit()

	// First call runs the initializer; subsequent calls return the cached value.
	fmt.Printf("%d\n", l.Value())
	fmt.Printf("%d\n", l.Value())
}
