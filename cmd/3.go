package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func main() {
	sum := 0
	wg := &sync.WaitGroup{}
	var mtx sync.Mutex
	//mtx := &sync.Mutex{}
	for i := 2; i <= 10; i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mtx.Lock()
			sum += i * i
			mtx.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(sum)
}
