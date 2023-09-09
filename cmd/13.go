package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("a:%d b:%d", a, b)
}
