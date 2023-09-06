package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
//произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
func worker(ch <-chan string, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for v := range ch {
		//time.Sleep(time.Millisecond)
		fmt.Printf("%d: %s\n", id, v)
	}
	fmt.Println(id, " finished")
}
func main() {
	nums := 0
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&nums)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		fmt.Println("Gracefully shutdown")
		cancel()
		time.Sleep(time.Second)
		close(ch)
		wg.Wait()
		os.Exit(0)
	}()

	for i := 0; i < nums; i++ {
		wg.Add(1)
		go worker(ch, wg, i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var val string
		for {
			fmt.Scan(&val)
			select {
			case <-ctx.Done():
				return
			default:
				ch <- val
			}

		}
	}()
	wg.Wait()
	fmt.Println("Finish")
}
