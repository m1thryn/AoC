package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var sid []int

	for s.Scan() {
		row := binSortX(128, s.Text()[:7])
		col := binSortX(8, s.Text()[7:])

		sid = append(sid, getSeatID(row, col))

	}
	var high int
	for _, val := range sid {
		if val > high {
			high = val
		}
	}
	fmt.Println("the highest SID is", high)
	getMySeatID(sid, high)
}

func binSortX(n int, s string) int {
	lower := 0
	upper := n - 1
	diff := n >> 1

	for _, v := range s {
		if v == 'F' || v == 'L' {
			upper -= diff
		}
		if v == 'B' || v == 'R' {
			lower += diff
		}
		diff >>= 1
	}
	if upper == lower {
		return upper
	}
	return 0
}

func getSeatID(r, c int) int {
	return (r * 8) + c
}

func getMySeatID(s []int, num int) {
	sort.Ints(s)

	for i := 100; i < num-100; i++ {
		if ok := find(s, i); !ok {
			fmt.Println("My ID is", i)
		}
	}
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
