// Solution to an Advent of Code problem, day 9, 2025
// https://adventofcode.com/2025/day/9
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

type Floor struct {
	RedTiles []RedTile `@@ (@@)*`
}

type RedTile struct {
	X int `@Int`
	Y int `","@Int`
}

type Interval struct {
	A int
	B int
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	parser, err := participle.Build[Floor]()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	floor, err := parser.Parse(inputFile, f)
	if err != nil {
		log.Fatal(err)
	}

	maxArea := 0
	for i := 0; i < len(floor.RedTiles); i++ {
		for j := i + 1; j < len(floor.RedTiles); j++ {
			dx := (floor.RedTiles[i].X - floor.RedTiles[j].X)
			dy := (floor.RedTiles[i].Y - floor.RedTiles[j].Y)
			dx = max(dx, -dx) + 1
			dy = max(dy, -dy) + 1
			area := dx * dy
			maxArea = max(maxArea, max(area, -area))
		}
	}

	fmt.Println("Part 1:", maxArea)

	N := 100000
	field := [100000][100000]byte{}
	for i := range N {
		for j := range N {
			field[i][j] = '-'
		}
	}

	for i := range floor.RedTiles {
		current := floor.RedTiles[i]
		next := floor.RedTiles[(i+1)%len(floor.RedTiles)]

		dy := (next.Y - current.Y)

		if dy > 0 {
			// flip left
			for y := min(current.Y, next.Y); y <= max(current.Y, next.Y); y++ {
				field[y][current.X] = '<'
			}

		} else if dy < 0 {
			// flip right
			for y := min(current.Y, next.Y); y <= max(current.Y, next.Y); y++ {
				field[y][current.X] = '>'
			}
		}

	}

	// paint
	intervals := [100000][]Interval{}
	for y := 0; y < N; y++ {
		in := false
		from := 0
		for x := 0; x < N; x++ {
			// Following assumes clock-wise polygon orientation. 
			// To generalize, non-zero winding number rule could be used instead of '<' and '>', where '>' and '<' are +-1.
			// https://www.tutorialspoint.com/computer_graphics/polygon_filling_algorithm.htm
			if field[y][x] == '<' && (x == N-1 || field[y][x+1] == '-') {
				in = false
				intervals[y] = append(intervals[y], Interval{from, x})
			} else if field[y][x] == '>' && !in {
				in = true
				from = x
			}
		}
	}

	maxArea = 0
	for i := 0; i < len(floor.RedTiles); i++ {
		for j := i + 1; j < len(floor.RedTiles); j++ {
			dx := (floor.RedTiles[i].X - floor.RedTiles[j].X)
			dy := (floor.RedTiles[i].Y - floor.RedTiles[j].Y)
			dx = max(dx, -dx) + 1
			dy = max(dy, -dy) + 1
			area := dx * dy
			if area < maxArea {
				continue
			}

			covered := true
			x0 := min(floor.RedTiles[i].X, floor.RedTiles[j].X)
			x1 := max(floor.RedTiles[i].X, floor.RedTiles[j].X)
			y0 := min(floor.RedTiles[i].Y, floor.RedTiles[j].Y)
			y1 := max(floor.RedTiles[i].Y, floor.RedTiles[j].Y)
			for y := y0; y <= y1; y++ {
				intervalFound := false
				for _, interval := range intervals[y] {
					intervalFound = intervalFound || (interval.A <= x0 && interval.B >= x1)
				}
				covered = covered && intervalFound
			}

			if covered {
				maxArea = max(maxArea, area)
			}
		}
	}

	fmt.Println("Part 2:", maxArea)
}
