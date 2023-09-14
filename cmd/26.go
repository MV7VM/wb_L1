package main

import "fmt"

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные,
//false etc). Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	m := make(map[rune]bool)
	//fmt.Printf("%d %d\n%d %d", 'A', 'Z', 'a', 'z')
	//65 90
	//97 122
	var s string
	fmt.Scan(&s)
	for _, v := range s {
		if v >= 65 && v <= 90 {
			v += 32
		}
		_, ok := m[v]
		m[v] = true
		if ok != false {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")

}
