package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	//1
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Stop_1")
				wg.Done()
				return
			default:
				fmt.Println("Run_1")
			}
		}
	}()
	time.Sleep(time.Millisecond)
	close(ch)
	//2
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stop_2")
				wg.Done()
				return
			default:
				fmt.Println("Run_2")
			}
		}
	}()
	time.Sleep(time.Millisecond)
	cancel()
	//3
	go func() {
		fmt.Println("stop_3")
		runtime.Goexit()
	}()
	wg.Wait()
}
