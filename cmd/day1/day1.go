package day1

import (
	"aoc2024/internal/util"
	"bufio"
	"fmt"
	"slices"
)

type Challenge struct {
	Scanner *bufio.Scanner
}

func (d *Challenge) Part1() {
	// 1. Parse input: list of two number pairs seperated by spaces.
	// 2. Sort both lists, ascending.
	// 3. Calculate abs diff of each pair and sum them.
	var list1, list2 []int
	for d.Scanner.Scan() {
		var n1, n2 int
		_, err := fmt.Sscanf(d.Scanner.Text(), "%d %d", &n1, &n2)
		if err != nil {
			util.ExitErr(1, "Error parsing input:", err)
		}
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	// Assert both lists are of equal length.
	if len(list1) != len(list2) {
		util.ExitErr(1, "Lists are not of equal length.")
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var sum int
	for i := 0; i < len(list1); i++ {
		n := list1[i] - list2[i]
		// Absolute value.
		if n < 0 {
			n = -n
		}
		sum += n
	}
	fmt.Println("Sum:", sum)
}

func (d *Challenge) Part2() {
	// 1. Parse input: instead of two lists, use a list for the first number and a map with a count for the second number.
	// 2. Multiply the numbers in the first list with the number of occurrences of the number in the second list, and
	// sum the results.
	var leftList []int
	var rightMap = make(map[int]int)
	for d.Scanner.Scan() {
		var n1, n2 int
		_, err := fmt.Sscanf(d.Scanner.Text(), "%d %d", &n1, &n2)
		if err != nil {
			util.ExitErr(1, "Error parsing input:", err)
		}
		leftList = append(leftList, n1)
		rightMap[n2]++
	}

	var sum int
	for _, n := range leftList {
		sum += n * rightMap[n]
	}

	fmt.Println("Sum:", sum)
}
