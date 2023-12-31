package main

import "fmt"

//-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5
//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func main() {
	m := make(map[float32][]float32)
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, v := range arr {
		id := float32(int(v) - int(v)%10.0)
		m[id] = append(m[id], v)
	}
	fmt.Println(m)
}
