package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	m := make(map[string]bool)
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range arr {
		m[v] = true
	}
	fmt.Println(m)
}
