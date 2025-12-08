// Solution to an Advent of Code problem, day 8, 2025
// https://adventofcode.com/2025/day/8
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Network struct {
	Boxes []Box `@@ (@@)*`
}

type Box struct {
	X int `@Int`
	Y int `","@Int`
	Z int `","@Int`
}

type Pair struct {
	A      int
	B      int
	DistSq int
}

func sortedPairs(boxes []Box) []Pair {
	pairs := []Pair{}
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			dx := boxes[i].X - boxes[j].X
			dy := boxes[i].Y - boxes[j].Y
			dz := boxes[i].Z - boxes[j].Z

			pairs = append(pairs, Pair{i, j, dx*dx + dy*dy + dz*dz})
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].DistSq < pairs[j].DistSq })
	return pairs
}

func dfs(color int, i int, colors []int, connections [][]int) int {
	if colors[i] != 0 {
		return 0
	}

	count := 1
	colors[i] = color
	for _, j := range connections[i] {
		count += dfs(color, j, colors, connections)
	}

	return count
}

func isFullyConnected(colors []int, connections [][]int) bool {
	for i := range colors {
		colors[i] = 0
	}

	countColored := dfs(1, 0, colors, connections)
	return countColored == len(colors)
}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading", inputFile)

	parser, err := participle.Build[Network]()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	network, err := parser.Parse(inputFile, f)
	if err != nil {
		log.Fatal(err)
	}
	// pp.Println(network)

	connections := make([][]int, len(network.Boxes))
	for i := range len(network.Boxes) {
		connections[i] = []int{}
	}

	pairs := sortedPairs(network.Boxes)

	nPairs := 1000
	if strings.Contains(inputFile, "test_input") {
		nPairs = 10
	}

	for pi := range nPairs {
		pair := pairs[pi]
		connections[pair.A] = append(connections[pair.A], pair.B)
		connections[pair.B] = append(connections[pair.B], pair.A)
		// pp.Println("Connect", network.Boxes[pair.A], network.Boxes[pair.B])
	}

	colors := make([]int, len(network.Boxes))
	nextColor := 1
	circuitSizes := []int{}
	for i := range colors {
		countColored := dfs(nextColor, i, colors, connections)
		if countColored > 0 {
			// fmt.Println(countColored)
			nextColor++
			circuitSizes = append(circuitSizes, countColored)
		}
	}

	sort.Slice(circuitSizes, func(i, j int) bool { return circuitSizes[j] < circuitSizes[i] })

	part1 := 1
	for i := range 3 {
		part1 *= circuitSizes[i]
	}

	fmt.Println("Part 1:", part1)

	for i := nPairs; i < len(pairs); i++ {
		pair := pairs[i]
		connections[pair.A] = append(connections[pair.A], pair.B)
		connections[pair.B] = append(connections[pair.B], pair.A)

		if isFullyConnected(colors, connections) {
			// pp.Println(network.Boxes[pair.A], network.Boxes[pair.B])
			fmt.Println("Part 2:", network.Boxes[pair.A].X*network.Boxes[pair.B].X)
			break
		}
	}
}
