package main

import "fmt"

func main() {
	var arr = []int{1, 9, 8, 7, 56, 6, 4, 3, 5, 8, 5, 32, 4}
	qsort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func qsort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	qsort(arr, start, pivot-1)
	qsort(arr, pivot+1, end)

}
func partition(arr []int, start, end int) int {
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < arr[end] {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
