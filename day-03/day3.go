package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const inputFile = "input1.txt"

type Position struct {
	x, y int
}

func (a *Position) Add(v *Position) {
	a.x += v.x
	a.y += v.y
}

var directions = map[string]Position{
	"v": {0, -1},
	"<": {-1, 0},
	">": {1, 0},
	"^": {0, 1},
}

func getInput() string {
	bs, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func walkHouses(houses string, grid map[Position]int) {
	curPos := Position{}
	grid[curPos] += 1
	for _, d := range houses {
		nextDirection := directions[string(d)]
		curPos.Add(&nextDirection)
		//fmt.Printf("%q [%d %d]\n", d, curPos.x, curPos.y)
		grid[curPos] += 1
	}
}

func splitPuzzle(houses string) (string, string) {
	var santa, robot string
	for i, c := range houses {
		if i%2 == 0 {
			santa += string(c)
		} else {
			robot += string(c)
		}
	}
	return santa, robot

}

func main() {
	puzzle := getInput()
	grid := make(map[Position]int)
	walkHouses(puzzle, grid)
	fmt.Printf("Part 1: %d\n", len(grid))

	grid2 := make(map[Position]int)
	santa, robot := splitPuzzle(puzzle)
	walkHouses(santa, grid2)
	walkHouses(robot, grid2)
	fmt.Printf("Part 2: %d\n", len(grid2))
}
