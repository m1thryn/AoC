package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	defer file.Close()

	var password []string
	var char []string
	var ranges []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ranges = append(ranges, strings.Fields(scanner.Text())[0])
		char = append(char, strings.Fields(scanner.Text())[1])
		password = append(password, strings.Fields(scanner.Text())[2])
	}

	// remove : from chars
	for i, val := range char {
		char[i] = val[:len(val)-1]
	}

	// goodPass, badPass := validatePasswords(password, char, ranges)
	goodPass, badPass := validatePositions(password, char, ranges)

	fmt.Println("Number of good passwords", goodPass)
	fmt.Println("Number of bad passwords", badPass)

}

func validatePasswords(a []string, b []string, c []string) (int, int) {
	var goodPass int
	var badPass int
	var count int = 0

	for i := range a {
		count = strings.Count(a[i], b[i])
		minVal, err := strconv.Atoi(strings.Split(c[i], "-")[0])
		if err != nil {
			log.Println("Error converting to string", err)
		}
		maxVal, err := strconv.Atoi(strings.Split(c[i], "-")[1])
		if err != nil {
			log.Println("Error converting to string", err)
		}

		if (count >= minVal) && (count <= maxVal) {
			goodPass++
		} else {
			badPass++
		}
	}

	return goodPass, badPass
}

func validatePositions(a []string, b []string, c []string) (int, int) {
	var goodPass int
	var badPass int

	for i := range a {
		pos1, err := strconv.Atoi(strings.Split(c[i], "-")[0])
		if err != nil {
			log.Println("Error converting to string", err)
		}
		pos2, err := strconv.Atoi(strings.Split(c[i], "-")[1])
		if err != nil {
			log.Println("Error converting to string", err)
		}

		if (string(a[i][pos1-1]) == b[i]) && (string(a[i][pos2-1]) == b[i]) {
			badPass++
		} else if (string(a[i][pos1-1]) == b[i]) || (string(a[i][pos2-1]) == b[i]) {
			goodPass++
		} else {
			badPass++
		}
	}

	return goodPass, badPass
}
