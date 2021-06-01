package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = "input1.txt"

type Command int

const (
	Toggle Command = iota
	On
	Off
	Unknown
)

type Instruction struct {
	command Command
	startX  int
	startY  int
	endX    int
	endY    int
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

func parse(puzzle []string) []Instruction {
	r := regexp.MustCompile(`([a-z ]+)(\d+),(\d+) through (\d+),(\d+)`)
	returnList := []Instruction{}
	for _, s := range puzzle {
		matches := r.FindStringSubmatch(s)
		var command Command
		command = Unknown
		switch {
		case strings.Contains(matches[1], "toggle"):
			command = Toggle
		case strings.Contains(matches[1], "turn on"):
			command = On
		case strings.Contains(matches[1], "turn off"):
			command = Off
		}
		fmt.Printf("%s \n", matches[1])
		startX, _ := strconv.Atoi(matches[2])
		startY, _ := strconv.Atoi(matches[3])
		endX, _ := strconv.Atoi(matches[4])
		endY, _ := strconv.Atoi(matches[5])
		returnList = append(returnList, Instruction{command, startX, startY, endX, endY})
	}

	return returnList
}

func mark(array *[1000][1000]bool, instr Instruction) {
	for i := instr.startX; i <= instr.endX; i++ {
		for j := instr.startY; j <= instr.endY; j++ {
			switch {
			case instr.command == Toggle:
				array[i][j] = !array[i][j]
			case instr.command == On:
				array[i][j] = true
			case instr.command == Off:
				array[i][j] = false
			}
		}
	}
}

func count(array *[1000][1000]bool) int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if array[i][j] {
				count++
			}
		}
	}
	return count
}

func main() {
	var twodim [1000][1000]bool
	theInput := getInput()
	parsed := parse(theInput)
	for _, c := range parsed {
		fmt.Printf("%+v\n", c)
		mark(&twodim, c)
	}
	fmt.Printf("%d\n", count(&twodim))
}
