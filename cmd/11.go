package main

import "fmt"

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for key := range set1 {
		if set2[key] {
			result[key] = true
		}
	}
	return result
}

func main() {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	set1[1] = true
	set1[2] = true
	set1[3] = true
	set2[2] = true
	set2[3] = true
	set2[4] = true
	result := intersection(set1, set2)
	fmt.Println("Пересечение множеств:", result)
}
