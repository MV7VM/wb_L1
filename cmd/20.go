package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow starts at 9"
	strArr := strings.Fields(s)
	for i := len(strArr) - 1; i >= 0; i-- {
		if i != 0 {
			fmt.Printf("%s ", strArr[i])
		} else {
			fmt.Printf("%s", strArr[i])
		}
	}
}
