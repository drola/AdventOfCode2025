// Solution to an Advent of Code problem, day 12, 2025
// https://adventofcode.com/2025/day/12
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for {
		var n int
		_, err := fmt.Fscanln(f, &n)
		if err != nil {
			break
		}
		sum += n
	}

	fmt.Println("Sum:", sum)
}
