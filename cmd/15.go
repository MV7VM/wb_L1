package main

import (
	"fmt"
	"unsafe"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = trim(v, 2, 100)
	fmt.Println(len(justString))
}

func createHugeString(i int) string {
	p := make([]rune, i)
	s := *(*string)(unsafe.Pointer(&p))
	return s
}
func trim(v string, from, to int) string {
	tmp := []rune(v)
	return string(append(tmp[from:from], tmp[from:to]...))
}
func main() {
	someFunc()
}
