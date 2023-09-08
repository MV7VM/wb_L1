package main

import "fmt"

func main() {
	var a int64
	a = 9
	fmt.Printf("%0b\n", a)
	changeBit(&a, 2, 1)
	fmt.Printf("%0b\n", a)
}

func changeBit(value *int64, position int, bit int) {
	if position > 64 {
		return
	}
	if bit == 0 {
		*value &^= 1 << position
	} else {
		*value |= 1 << position
	}
}
