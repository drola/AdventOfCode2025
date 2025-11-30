#!/bin/bash

echo -n "Please enter the day number (example: 11): "
read day_number
if [ -z "$day_number" ]; then
    echo "Day number is required."
    exit 1
fi

day_number_padded=$(printf '%02d' ${day_number})
mkdir -p "./day${day_number_padded}"
tee "./day${day_number_padded}/main.go" <<EOF >/dev/null
// Solution to an Advent of Code problem, day ${day_number}, 2025
// https://adventofcode.com/2025/day/${day_number}
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
EOF

touch "./day${day_number_padded}/test_input.txt"
touch "./day${day_number_padded}/input.txt"
