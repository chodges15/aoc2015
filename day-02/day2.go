package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = "input1.txt"

const dimensions = 3

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

func parseInput(puzzle []string) [][]int {
	length := len(puzzle)
	rVal := make([][]int, length)
	for i, dim := range puzzle {
		rVal[i] = make([]int, dimensions)
		for j, num := range strings.Split(dim, "x") {
			con, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			} else {
				rVal[i][j] = con
			}
		}
		sort.Ints(rVal[i])
	}
	return rVal
}

func wrappingPaperNeeded(sortedInput [][]int) int {
	area := 0
	for _, gift := range sortedInput {
		area += 2 * gift[0] * gift[1] + 2 * gift[1] * gift[2] + 2 * gift[0] * gift[2]
		area += gift[0] * gift[1]
	}
	return area
}

func main() {
	sortedInput := parseInput(getInput())
	fmt.Println(wrappingPaperNeeded(sortedInput))
}
