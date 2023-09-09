package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func main() {
	a := 123.
	detectType(a)
}
func detectType(a interface{}) {
	switch v := a.(type) {
	default:
		fmt.Printf("%T", v)
	}
}
