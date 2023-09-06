package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	var a = [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, val := range a {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			powVal := val * val
			fmt.Println(powVal)
		}(val)
	}
	wg.Wait()
}
