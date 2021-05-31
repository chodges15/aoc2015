package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "input1.txt"

var forbidden = []string{
	"ab",
	"cd",
	"pq",
	"xy",
}

func getInput() []string {
	returnList := []string{}
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		returnList = append(returnList, scanner.Text())
	}

	return returnList
}

func niceString(str string) bool {
	doubleLetter := false
	lastLetter := ""
	for _, f := range forbidden {
		if strings.Contains(str, f) {
			return false
		}
	}
	for _, c := range str {
		if lastLetter == string(c) {
			doubleLetter = true
		}
		lastLetter = string(c)
	}
	vowelsCount := strings.Count(str, "a") + strings.Count(str, "e") + strings.Count(str, "i") + strings.Count(str, "o") + strings.Count(str, "u")
	return doubleLetter && vowelsCount >= 3
}

func niceString2(str string) bool {
	pairTwice := false
	oneLetterBetween := false
	for i := 2; i < len(str); i++ {
		pair := str[i-2 : i]
		if strings.Contains(str[i:], pair) {
			pairTwice = true
			fmt.Println("PairTwice")
			break
		}
	}
	if !pairTwice {
		return false
	}

	for i := 0; i < len(str) - 2; i++ {
		if str[i] == str[i+2] {
			oneLetterBetween = true
			fmt.Println("OneLetterBetween")
			break
		}
	}

	return pairTwice && oneLetterBetween
}

func main() {
	puzzle := getInput()
	var count1, count2 int
	for _, str := range puzzle {
		fmt.Println(str)
		if niceString(str) {
			count1++
		}
		if niceString2(str) {
			count2++
		}
	}
	fmt.Printf("Part 1: %d\n", count1)
	fmt.Printf("Part 2: %d\n", count2)
}
