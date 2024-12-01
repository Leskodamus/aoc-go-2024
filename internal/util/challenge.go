package util

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Challenge interface {
	Part1()
	Part2()
}

func ExitErr(errCode int, a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	flag.Usage()
	os.Exit(errCode)
}

func ReadInput(input string) (*bufio.Scanner, error) {
	if input == "" {
		return nil, fmt.Errorf("no input file specified.")
	}

	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %s", err)
	}

	return bufio.NewScanner(file), nil
}
