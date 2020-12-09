package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fContent := string(f)

	groups := strings.Split(fContent, "\n\n")

	fmt.Println("Number of unique chars", getUniqueChars2(groups))

}

func getUniqueChars(s []string) int {
	var total int

	for _, v := range s {
		replacer := strings.NewReplacer("\n", "")
		v := replacer.Replace(v)

		chars := make(map[rune]string)

		for _, v := range v {
			if _, ok := chars[v]; !ok {
				chars[v] = ""
			}
		}
		total += len(chars)
	}
	return total
}

//cheater.. :-(
func getUniqueChars2(groups []string) int {
	var answerSum int
	for _, group := range groups {
		persons := strings.Split(group, "\n")

		answers := make(map[rune]bool)
		for _, char := range persons[0] {
			answers[char] = true
		}
		for _, person := range persons[1:] {
			newAnswers := make(map[rune]bool)
			for _, char := range person {
				if _, ok := answers[char]; ok {
					newAnswers[char] = true
				}
			}
			answers = newAnswers
		}

		answerSum += len(answers)
	}

	return answerSum
}
