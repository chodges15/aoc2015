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

func main() {
	puzzle := getInput()
	var count int
	for _, str := range puzzle {
		fmt.Println(str)
		if niceString(str) {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}
