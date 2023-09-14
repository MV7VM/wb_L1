package main

import (
	"L1/cmd/24/point"
	"fmt"
	"math"
)

func main() {
	a := point.New()
	b := point.New()
	a.SetX(0)
	a.SetY(0)
	b.SetX(1)
	b.SetY(1)
	fmt.Println(distance(a.GetX(), a.GetY(), b.GetX(), b.GetY()))
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
