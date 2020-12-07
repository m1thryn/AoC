package main

import (
	"fmt"
)

func main() {
	var b []int
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 7, 8, 9, 10)

	for i := 0; i < 2; i++ {
		b = append(b, a...)
	}
	fmt.Println(a)
	fmt.Println(b)
}
