package main

import (
	"fmt"
	"math/rand"
)

func binSearch(start, end, ans int) int {
	mid := start + (end-start)/2
	for {
		fmt.Println(start, end, ans)
		mid = start + (end-start)/2
		if start >= end {
			return start
		}
		if mid > ans {
			end = mid
		} else if mid < ans {
			start = mid
		} else {
			return mid
		}
	}
	return mid
}

func main() {
	ans := rand.Int() % 10240
	fmt.Printf("ans is %d and bin search ans is %d\n", ans, binSearch(0, 10240, ans))
}
