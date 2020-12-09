package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var a []string

	for s.Scan() {
		var extendedLine string
		for i := 0; i < 100; i++ {
			extendedLine += s.Text()
		}
		a = append(a, extendedLine)
	}
	total := 1
	total *= getTrees(a, 1, 1)
	total *= getTrees(a, 3, 1)
	total *= getTrees(a, 5, 1)
	total *= getTrees(a, 7, 1)
	total *= getTrees(a, 1, 2)

	fmt.Println(total)
}

func getTrees(s []string, x, y int) int {
	treeCount := 0
	xPos := 0
	yPos := 0

	for yPos < len(s) {
		if s[yPos][xPos] == '#' {
			treeCount++
		}
		xPos += x
		yPos += y
	}
	return treeCount

}
