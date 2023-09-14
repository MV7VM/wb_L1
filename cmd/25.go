package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	select {
	case <-time.After(duration):
	}
}

func main() {
	duration := time.Second * 2
	fmt.Println("here")
	mySleep(duration)
	fmt.Println("there")
}
