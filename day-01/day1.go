package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = "input1.txt"

func getInput() string {
	bs, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	puzzle := getInput()
	fmt.Printf("Part 1 %d\n", part1(puzzle))
	fmt.Printf("Part 2 %d\n", part2(puzzle))
}

func part1(puzzle string) int {
	return strings.Count(puzzle, "(") - strings.Count(puzzle, ")")
}

func part2(puzzle string) int {
	currentFloor := 0
	for i, c := range puzzle {
		if c == '(' {
			currentFloor++
		} else {
			currentFloor--
		}
		if currentFloor < 0 {
			return i + 1
		}
	}
	return -1
}
