package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		dummy, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Couldn't convert string to int:", err)
		}

		nums = append(nums, dummy)
	}
	// sol1, sol2 := getTwoSum(nums, 2020)
	// fmt.Println(sol1, sol2)
	// fmt.Println(sol1 * sol2)

	sol1, sol2, sol3 := getThreeSum(nums, 2020)
	fmt.Println(sol1, sol2, sol3)
	fmt.Println(sol1 * sol2 * sol3)

}

func getTwoSum(nums []int, goal int) (int, int) {

	for _, val := range nums {
		findMe := goal - val

		if contains(nums, findMe) {
			return val, findMe
		}

	}
	return 0, 0
}

func getThreeSum(nums []int, goal int) (int, int, int) {

	for _, val1 := range nums {
		for _, val2 := range nums {

			findMe := goal - (val1 + val2)
			if contains(nums, findMe) {
				return val1, val2, findMe
			}
		}

	}
	return 0, 0, 0
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
