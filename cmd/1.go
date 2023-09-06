package main

import "fmt"

type Human struct {
	a int
	b int
}

type Action struct {
	Human
}

func (h *Human) sum() int {
	return h.a + h.b
}

// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования).
func main() {
	var A Action = Action{Human{1, 2}}
	fmt.Println(A.sum())
	A.a = 3
	fmt.Println(A.sum())

}
