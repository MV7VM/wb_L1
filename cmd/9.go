package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := make(chan int)
	b := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer close(a)
		defer wg.Done()
		for i := 0; i < len(arr); i++ {
			a <- arr[i]
		}
		return
	}()
	go func() {
		defer close(b)
		defer wg.Done()
		val := 0
		for i := 0; i < len(arr); i++ {
			val = <-a
			b <- val * 2
		}
		return
	}()
	go func() {
		defer wg.Done()
		for v := range b {
			fmt.Println(v)
		}
		return
	}()
	wg.Wait()
}
