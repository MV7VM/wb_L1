package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	mu    sync.Mutex
}

func (c *counter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	wg := &sync.WaitGroup{}
	c := &counter{value: 0, mu: sync.Mutex{}}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}
