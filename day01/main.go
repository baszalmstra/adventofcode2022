package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	// Parse the input
	elves := [][]int{}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	for _, group := range lines {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			num, _ := strconv.Atoi(line)
			row = append(row, num)
		}
		elves = append(elves, row)
	}

	// Compute total per elf
	elfTotals := []int{}
	for _, items := range elves {
		total := 0
		for _, cals := range items {
			total += cals
		}
		elfTotals = append(elfTotals, total)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotals)))

	println("Solution 1:", elfTotals[0])
	println("Solution 2:", elfTotals[0]+elfTotals[1]+elfTotals[2])
}
