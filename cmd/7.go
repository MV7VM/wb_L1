package main

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// реализовать конкурентную запить в map

func main() {
	var id int
	m := make(map[int]int)
	mtx := &sync.Mutex{}
	ctx, cancel := context.WithCancel(context.Background())
	new_id := func() int {
		id++
		return id
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				mtx.Lock()
				m[new_id()] = rand.Int()
				mtx.Unlock()
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				mtx.Lock()
				m[new_id()] = rand.Int()
				mtx.Unlock()
			}
		}
	}()
	time.Sleep(20 * time.Microsecond)
	cancel()
	wg.Wait()
	arr := make([]int, 1)
	for k, _ := range m {
		arr = append(arr, k)
		//fmt.Printf(" %d ", k)
	}
	sort.Ints(arr)
	fmt.Println(arr)
}
