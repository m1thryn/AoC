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
		log.Fatalf("Error opening input file: %s", err)
	}

	rawLines := parseLines(file)

	var goodPassports int

	for _, val := range rawLines {
		if len(val) == 8 {
			if validateFields(val) {
				goodPassports++
			}
		} else if len(val) == 7 {
			if _, ok := val["cid"]; !ok {
				if validateFields(val) {
					goodPassports++
				}
			}
		}
	}

	fmt.Println(goodPassports)
}

func parseLines(f *os.File) []map[string]string {
	var passports []map[string]string
	var passport string
	var rawLine []string
	s := bufio.NewScanner(f)

	for s.Scan() {

		if len(s.Text()) != 0 {
			passport += s.Text() + " "
		} else {
			rawLine = append(rawLine, passport)
			passport = ""
		}

	}

	for _, val := range rawLine {
		m := newMap()
		for _, val := range strings.Fields(val) {
			items := strings.Split(val, ":")
			m[items[0]] = items[1]
		}
		passports = append(passports, m)
	}

	return passports
}

func newMap() map[string]string {
	m := make(map[string]string)
	return m
}

func validateFields(m map[string]string) bool {
	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	var token int

	if val, _ := strconv.Atoi(m["byr"]); val >= 1920 && val <= 2002 {
		token++
	}
	if val, _ := strconv.Atoi(m["iyr"]); val >= 2010 && val <= 2020 {
		token++
	}
	if val, _ := strconv.Atoi(m["eyr"]); val >= 2020 && val <= 2030 {
		token++
	}
	if m["hgt"][len(m["hgt"])-2:] == "cm" {
		if val, _ := strconv.Atoi(m["hgt"][:len(m["hgt"])-2]); val >= 150 && val <= 193 {
			token++
		}
	}
	if m["hgt"][len(m["hgt"])-2:] == "in" {
		if val, _ := strconv.Atoi(m["hgt"][:len(m["hgt"])-2]); val >= 59 && val <= 76 {
			token++
		}
	}
	if val := m["hcl"]; string(val[0]) == "#" {
		if len(val) == 7 {
			var count int
			for _, v := range val[1:7] {
				if find(hex, string(v)) {
					count++
				}
			}
			if count == 6 {
				token++
			}
		}
	}
	if find(eyeColors, m["ecl"]) {
		token++
	}
	if val := m["pid"]; len(val) == 9 {
		token++
	}

	fmt.Println(token)
	if token == 7 {
		return true
	}
	return false
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
