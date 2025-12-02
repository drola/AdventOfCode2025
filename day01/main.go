// Solution to an Advent of Code problem, day 1, 2025
// https://adventofcode.com/2025/day/1
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	dialPosition := 50
	finalZerosCount := 0
	allZerosCount := 0
	for {
		var line string
		_, err := fmt.Fscanln(f, &line)
		if err != nil {
			break
		}

		if len(line) > 1 {
			direction := line[0]
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal(err)
			}
			if direction == 'L' {
				fmt.Println("-", steps)
				for i := 0; i < steps; i++ {
					dialPosition = (100 + dialPosition - 1) % 100
					if dialPosition == 0 {
						allZerosCount++
					}

				}
			} else {
				fmt.Println("+", steps)
				for i := 0; i < steps; i++ {
					dialPosition = (100 + dialPosition + 1) % 100
					if dialPosition == 0 {
						allZerosCount++
					}

				}
			}

			fmt.Println(dialPosition)

			if dialPosition == 0 {
				finalZerosCount++
			}

		}
	}

	fmt.Println("Zeros count:", finalZerosCount)
	fmt.Println("All Zeros count:", allZerosCount)
}
