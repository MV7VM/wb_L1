package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func main() {
	var nums string
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Print("ENter the time for stop: ")
	fmt.Scan(&nums)
	duration, _ := time.ParseDuration(nums)
	timer := time.NewTimer(duration)
	//IsTime := make(chan bool)
	//wg.Add(1)
	//go func() {
	//	<-sig
	//	fmt.Println("Gracefully shutdown")
	//	cancel()
	//	time.Sleep(time.Second)
	//	close(ch)
	//	wg.Wait()
	//	os.Exit(0)
	//}()

	go func() {
		//defer wg.Done()
		<-timer.C
		cancel()
		time.Sleep(150 * time.Millisecond)
		close(ch)
		wg.Wait()
		os.Exit(0)
		//IsTime <- true
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var val int
		for {
			val = rand.Int()
			select {
			case <-ctx.Done():
				return
			default:
				ch <- val
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			//time.Sleep(time.Millisecond)
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-ctx.Done():
				return
			}

		}
	}()
	wg.Wait()

}
