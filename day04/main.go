// Solution to an Advent of Code problem, day 4, 2025
// https://adventofcode.com/2025/day/4
package main

import (
	"fmt"
	"log"
	"os"
)

func isRoll(grid [][]byte, x int, y int) bool {
	if y >= len(grid) || y < 0 || x < 0 || x >= len(grid[y]) {
		return false
	}

	return grid[y][x] == '@'
}

func isAccessibleRoll(grid [][]byte, x int, y int) bool {
	if grid[y][x] == '@' {
		adjacentRolls := 0
		if isRoll(grid, x, y-1) {
			adjacentRolls++
		}
		if isRoll(grid, x, y+1) {
			adjacentRolls++
		}
		if isRoll(grid, x-1, y) {
			adjacentRolls++
		}
		if isRoll(grid, x+1, y) {
			adjacentRolls++
		}
		if isRoll(grid, x-1, y-1) {
			adjacentRolls++
		}
		if isRoll(grid, x-1, y+1) {
			adjacentRolls++
		}
		if isRoll(grid, x+1, y-1) {
			adjacentRolls++
		}
		if isRoll(grid, x+1, y+1) {
			adjacentRolls++
		}

		if adjacentRolls < 4 {
			return true
		}
	}

	return false
}

type Point struct {
	x int
	y int
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	grid := [][]byte{}

	for {
		var line []byte
		_, err := fmt.Fscanln(f, &line)
		if err != nil {
			break
		}
		if len(line) > 0 {
			grid = append(grid, line)
		}
	}

	w := len(grid[0])
	h := len(grid)

	accessibleRolls := 0
	for x := range w {
		for y := range h {
			if isAccessibleRoll(grid, x, y) {
				accessibleRolls++
			}
		}
	}

	fmt.Println("Accessible rolls [part 1]:", accessibleRolls)

	removedRolls := 0
	for {
		removalPlan := []Point{}
		for x := range w {
			for y := range h {
				if isAccessibleRoll(grid, x, y) {
					removalPlan = append(removalPlan, Point{x, y})
				}
			}
		}

		if len(removalPlan) == 0 {
			break
		}
		removedRolls += len(removalPlan)

		for _, p := range removalPlan {
			grid[p.y][p.x] = '.'
		}

	}
	fmt.Println("Removed rolls [part 2]:", removedRolls)

}
