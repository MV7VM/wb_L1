package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		a = append(a, rand.Int()%100)
	}
	id := rand.Int() % 10
	fmt.Printf("del %d from %d\n", id, a)
	a = append(a[:id], a[id+1:]...)
	fmt.Println("len = ", len(a), " arr: ", a)
}
