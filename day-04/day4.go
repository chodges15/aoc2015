package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
)

const inputFile = "input1.txt"

func getInput() string {
	bs, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

// not portable, but gets the job done on ARM
func acceptableHash(hash [16]byte, part int) bool {
	if part == 1 {
		return hash[0] == 0 && hash[1] == 0 && (hash[2]&0xf0) == 0
	} else {
		return hash[0] == 0 && hash[1] == 0 && hash[2] == 0
	}
}


func findAcceptableHash(puzzle string, part int) int {
	hashFound := false
	fmt.Println(puzzle)
	answer := 0
	for n := 0; !hashFound; n++ {
		data := fmt.Sprintf("%s%d", puzzle, n)
		hash := md5.Sum([]byte(data))
		hashFound = acceptableHash(hash, part)
		if hashFound {
			answer = n
		}
	}
	return answer

}

func main() {
	puzzle := getInput()
	fmt.Printf("Part 1: %d\n", findAcceptableHash(puzzle, 1))
	fmt.Printf("Part 2: %d\n", findAcceptableHash(puzzle, 2))
}
