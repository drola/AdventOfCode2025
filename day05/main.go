// Solution to an Advent of Code problem, day 5, 2025
// https://adventofcode.com/2025/day/5
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

type Input struct {
	Ranges      []*Range "@@ ( @@)*"
	Ingredients []int    `@Int ( @Int)*`
}

type Range struct {
	From int `@Int`
	To   int `"-"@Int`
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	parser, err := participle.Build[Input]()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input, err := parser.Parse(inputFile, f)
	if err != nil {
		log.Fatal(err)
	}
	// pp.Print(input)

	n := 0
	for _, ingredient := range input.Ingredients {
		for _, r := range input.Ranges {
			if ingredient >= r.From && ingredient <= r.To {
				n++
				break
			}
		}
	}

	fmt.Println("Part 1:", n)

	ranges := []*Range{}
	// clone ranges into array
	for _, r := range input.Ranges {
		ranges = append(ranges, &Range{r.From, r.To})
	}

	// keep shortening until no overlaps are found
	for {
		wasOverlap := false

		for i := 0; i < len(ranges); i++ {
			if ranges[i] == nil {
				continue
			}

			for j := i + 1; j < len(ranges); j++ {
				if ranges[j] == nil {
					continue
				}

				left := ranges[i]
				right := ranges[j]
				if right.From < left.From {
					left = ranges[j]
					right = ranges[i]
				}

				if left.To >= right.From { // overlap
					ranges[i].From = left.From
					ranges[i].To = max(left.To, right.To)
					ranges[j] = nil
					wasOverlap = true
				}

				if wasOverlap {
					break
				}
			}
			if wasOverlap {
				break
			}
		}

		if !wasOverlap {
			break
		}
	}

	acc := 0
	for _, r := range ranges {
		if r != nil {
			acc += r.To - r.From + 1
		}
	}
	fmt.Println("Part 2:", acc)
}
