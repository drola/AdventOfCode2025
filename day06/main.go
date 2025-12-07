// Solution to an Advent of Code problem, day 6, 2025
// https://adventofcode.com/2025/day/6
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/k0kubun/pp/v3"
)

type Numbers struct {
	Numbers []int `@Int ( @Int )*`
}

type Ops struct {
	Ops []string `@("+" | "*") ( @("+" | "*"))*`
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	numbers := [][]int{}
	operations := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		numbersParser, err := participle.Build[Numbers]()
		if err != nil {
			log.Fatal(err)
		}

		opsParser, err := participle.Build[Ops]()
		if err != nil {
			log.Fatal(err)
		}

		if strings.ContainsAny(line, "*+") {
			ops, err := opsParser.ParseString("input", line)
			if err != nil {
				log.Fatal(err)
			}

			operations = ops.Ops

		} else {
			nrs, err := numbersParser.ParseString("input", line)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, nrs.Numbers)
		}
	}

	grandTotal := 0
	for column, op := range operations {
		if op == "+" {
			r := 0
			for _, ln := range numbers {
				r += ln[column]
			}
			grandTotal += r
		}
		if op == "*" {
			r := 1
			for _, ln := range numbers {
				r *= ln[column]
			}
			grandTotal += r
		}
	}
	fmt.Println("Part 1:", grandTotal)

	f, err = os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	pp.Print(lines)

	op := ' '
	innerAcc := 0
	grandTotal = 0
	for x := 0; x < len(lines[0]); x++ {
		if lines[len(lines)-1][x] == '+' {
			grandTotal += innerAcc
			innerAcc = 0
			op = '+'
			fmt.Println("Set op to +")
		} else if lines[len(lines)-1][x] == '*' {
			grandTotal += innerAcc
			innerAcc = 1
			op = '*'
			fmt.Println("Set op to *")
		}

		operand := 0
		for y := 0; y < len(lines)-1; y++ {
			if lines[y][x] != ' ' {
				operand = operand*10 + int(lines[y][x]-'0')
			}
		}
		fmt.Println(operand)

		if operand > 0 {
			switch op {
			case '+':
				innerAcc += operand
			case '*':
				innerAcc *= operand
			}
		}
	}

	grandTotal += innerAcc
	fmt.Println("Part 2:", grandTotal)
}
