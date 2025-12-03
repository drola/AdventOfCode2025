// Solution to an Advent of Code problem, day 2, 2025
// https://adventofcode.com/2025/day/2
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alecthomas/participle/v2"
)

type Ranges struct {
	Ranges []*Range `@@ ("," @@)*`
}

type Range struct {
	From int `@Int`
	To   int `"-"@Int`
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	parser, err := participle.Build[Ranges]()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ranges, err := parser.Parse(inputFile, f)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, r := range ranges.Ranges {
		for id := r.From; id <= r.To; id++ {
			digits := strconv.Itoa(id)
			if len(digits)%2 != 0 {
				continue
			}
			invalid := true
			for idx := 0; idx < (len(digits) / 2); idx++ {
				invalid = invalid && digits[idx] == digits[idx+len(digits)/2]
			}
			if invalid {
				sum += id
			}
		}
	}

	fmt.Println("Parts 1:", sum)

	sum = 0
	for _, r := range ranges.Ranges {
		for id := r.From; id <= r.To; id++ {
			digits := strconv.Itoa(id)

			invalid := false
			for freq := 2; freq <= len(digits); freq++ {
				if len(digits)%freq != 0 {
					continue
				}
				invalid_for_freq := true
				for idx := 0; idx < (len(digits) / freq); idx++ {
					all_equal := true
					for rep := 0; rep < freq; rep++ {
						all_equal = all_equal && digits[idx] == digits[idx+rep*len(digits)/freq]
					}
					invalid_for_freq = invalid_for_freq && all_equal
				}
				invalid = invalid || invalid_for_freq
			}

			if invalid {
				sum += id
			}
		}
	}
	fmt.Println("Parts 2:", sum)
}
