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
			area := (floor.RedTiles[i].X + 1 - floor.RedTiles[j].X) * (floor.RedTiles[i].Y + 1 - floor.RedTiles[j].Y)
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

	// window := 15
	// for i := range window {
	// 	fmt.Println(string(field[i][0:window]))
	// }

	for i := range floor.RedTiles {
		current := floor.RedTiles[i]
		next := floor.RedTiles[(i+1)%len(floor.RedTiles)]
		// next2 := floor.RedTiles[(i+2)%len(floor.RedTiles)]

		dx := (next.X - current.X)
		// dx2 := (next2.X - next.X)
		dy := (next.Y - current.Y)
		// dy2 := (next2.Y - next.Y)

		if dy > 0 { //&& dx2 < 0 {
			// flip left
			for y := min(current.Y, next.Y); y <= max(current.Y, next.Y); y++ {
				if field[y][current.X] == '-' {
					field[y][current.X] = '<'
				}
			}

		} else if dy < 0 {
			// flip right
			for y := min(current.Y, next.Y); y <= max(current.Y, next.Y); y++ {
				// if field[y][current.X] == '-' {
				field[y][current.X] = '>'
				// }
			}
		} else if dx != 0 {
			for x := min(current.X, next.X); x <= max(current.X, next.X); x++ {
				if field[current.Y][x] == '-' {
					// field[current.Y][x] = 'X'
				}

			}

		}

		// pp.Println(current, next)
		// for i := range window {
		// 	fmt.Println(string(field[i][0:window]))
		// }
	}

	// paint

	intervals := [100000][]Interval{}
	for y := 0; y < N; y++ {
		in := false
		from := 0
		for x := 0; x < N; x++ {
			if field[y][x] == '<' && (x == N-1 || field[y][x+1] == '-') {
				in = false
				intervals[y] = append(intervals[y], Interval{from, x})
				field[y][x] = 'X'
			} else if field[y][x] == '>' && !in {
				in = true
				from = x
			}

			if in {
				field[y][x] = 'X'
			}
		}
	}
	// for i := range window {
	// 	fmt.Println(string(field[i][0:window]))
	// }

	maxArea = 0
	for i := 0; i < len(floor.RedTiles); i++ {
		for j := i + 1; j < len(floor.RedTiles); j++ {
			area := (floor.RedTiles[i].X + 1 - floor.RedTiles[j].X) * (floor.RedTiles[i].Y + 1 - floor.RedTiles[j].Y)
			area = max(area, -area)
			if area < maxArea {
				continue
			}

			covered := true
			x0 := min(floor.RedTiles[i].X, floor.RedTiles[j].X)
			x1 := max(floor.RedTiles[i].X, floor.RedTiles[j].X)
			y0 := min(floor.RedTiles[i].Y, floor.RedTiles[j].Y)
			y1 := max(floor.RedTiles[i].Y, floor.RedTiles[j].Y)
			for y := y0; y <= y1; y++ {
				for x := x0; x <= x1; x++ {
					covered = covered && field[y][x] != '-'
					if !covered {
						break
					}
				}
				if !covered {
					break
				}
				// intervalFound := false
				// for _, interval := range intervals[y] {
				// 	intervalFound = intervalFound || (interval.A <= x0 && interval.B >= x1)
				// }
				// covered = covered && intervalFound
			}

			if covered {
				maxArea = max(maxArea, area)
			}
		}
	}

	fmt.Println("Part 2:", maxArea)
	// 1474444748 is too low
	// 138584360 is too low
	// 314679376

	// pp.Println(intervals[0:15])

}
