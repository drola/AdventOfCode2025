// Solution to an Advent of Code problem, day 3, 2025
// https://adventofcode.com/2025/day/3
package main

import (
	"fmt"
	"log"
	"os"
)

func maxDigit(s string) (int, int) {
	resDigit := byte(0)
	resIdx := 0

	for i := 0; i < len(s); i++ {
		if s[i] > resDigit {
			resDigit = s[i]
			resIdx = i
		}
		if resDigit == '9' {
			break
		}
	}

	return int(resDigit - '0'), resIdx
}

func maxJoltage(line string, nDigits int) int {
	left := 0
	accumulator := 0
	for i := range nDigits {
		digit, digitIdx := maxDigit(line[left : len(line)-(nDigits-i-1)])
		left += digitIdx + 1
		// fmt.Println("Digit:", digit, "DigitIdx", digitIdx, "Left:", left)
		accumulator *= 10
		accumulator += digit
	}

	return accumulator
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	part1 := 0
	part2 := 0
	for {
		var line string
		_, err := fmt.Fscanln(f, &line)
		if err != nil {
			break
		}

		part1 += maxJoltage(line, 2)
		part2 += maxJoltage(line, 12)

		// fmt.Println(maxJoltage(line, 12))
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
