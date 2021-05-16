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

// not super portable, but gets the job done on ARM
func acceptableHash(hash [16]byte) bool {
	return hash[0] == 0 && hash[1] == 0 && (hash[2]&0xf0) == 0
}

func main() {
	puzzle := getInput()
	hashFound := false
	fmt.Println(puzzle)
	answer1 := 0
	for n := 0; !hashFound; n++ {
		data := fmt.Sprintf("%s%d", puzzle, n)
		hash := md5.Sum([]byte(data))
		// hashString := fmt.Sprintf("%x", hash)
		// fmt.Println(hashString)
		// fmt.Println(hashFound)
		hashFound = acceptableHash(hash)
		if hashFound {
			answer1 = n
		}
	}
	fmt.Printf("Part 1: %d\n", answer1)
}
