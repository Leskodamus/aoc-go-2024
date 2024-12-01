package main

/*
CLI syntax:
    go run main.go <day> <part> [input file]

    <day> - day of the challenge {1..25}
    <part> - part of the challenge {1,2}
    [input file] - optional input file, defaults to input.txt in directory of <day>
*/

import (
	"aoc2024/cmd/day1"
	"aoc2024/internal/util"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <day> <part> [input file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Parse CLI arguments
	var day, part int
	var input string

	flag.IntVar(&day, "day", 0, "day of the challenge {1..25}")
	flag.IntVar(&part, "part", 0, "part of the challenge {1,2}")
	flag.StringVar(&input, "input", "", "input txt file")

	flag.Parse()

	if flag.NArg() == 0 && day == 0 && part == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if !((day > 0 && day < 26) && (part == 1 || part == 2)) {
		util.ExitErr(1, "Invalid arguments: day", day, "part", part, "\n")
	}

	if input == "" {
		input = fmt.Sprintf("./cmd/day%d/input.txt", day)
	}

	scanner, err := util.ReadInput(input)
	if err != nil {
		util.ExitErr(1, err)
	}

	fmt.Println("===========================\n"+
		"Runnning challenge for day", day, "part", part, "\n"+
		"===========================\n",
	)

	var challenge util.Challenge

	switch day {
	case 1:
		challenge = &day1.Challenge{Scanner: scanner}
	default:
		util.ExitErr(1, "Day", day, "not implemented yet")
	}

	if part == 1 {
		challenge.Part1()
	} else {
		challenge.Part2()
	}
}
