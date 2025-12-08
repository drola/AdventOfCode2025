// Solution to an Advent of Code problem, day 7, 2025
// https://adventofcode.com/2025/day/7
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countTimelines(matrix []string, memory [][]int, x int, y int) int {

	if y == len(matrix) {
		return 1
	}
	if x < 0 || x >= len(matrix[0]) {
		return 0
	}
	if memory[y][x] > 0 {
		return memory[y][x]
	}

	if matrix[y][x] == '^' {
		result := countTimelines(matrix, memory, x-1, y+1) + countTimelines(matrix, memory, x+1, y+1)
		memory[y][x] = result
		return result
	} else {
		result := countTimelines(matrix, memory, x, y+1)
		memory[y][x] = result
		return result
	}
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	matrix := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			matrix = append(matrix, line)
		}
	}

	w := len(matrix[0])
	before := []byte(strings.ReplaceAll(matrix[0], "S", "|"))
	splitCount := 0
	fmt.Println(string(before))
	for i := 1; i < len(matrix); i++ {
		beams := make([]byte, w)
		for j := 0; j < w; j++ {
			if before[j] != '|' {
				continue
			}

			if matrix[i][j] == '^' {
				splitCount += 1
				if j > 0 {
					beams[j-1] = '|'
				}
				if j < w-1 {
					beams[j+1] = '|'
				}

			} else {
				beams[j] = '|'
			}

		}
		before = beams
		fmt.Println(string(before))
	}

	memory := make([][]int, len(matrix))
	for y := 0; y < len(matrix); y++ {
		memory[y] = make([]int, w)
	}
	fmt.Println("Part 1:", splitCount)
	for x := 0; x < w; x++ {
		if matrix[0][x] == 'S' {
			fmt.Println("Part 2:", countTimelines(matrix, memory, x, 0))
		}
	}

}
