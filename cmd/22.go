package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)
	a.SetInt64(4194304)
	b.SetInt64(4194304)
	//for i := 0; i < 22; i++ {
	//	a *= 2
	//	b *= 2
	//}
	res.Mul(a, b)
	fmt.Println(res)
	res.Div(a, b)
	fmt.Println(res)
	res.Add(a, b)
	fmt.Println(res)
	res.Sub(a, b)
	fmt.Println(res)
}
