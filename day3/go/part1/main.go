package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// read input from file and convert to string
	inputBytes, err := os.ReadFile("../../input.txt")
	if err != nil {
		panic(err)
	}
	corruptedMemory := string(inputBytes)

	// find all matches for valid mul command, include selector groups in output
	mulExpression := regexp.MustCompile(`mul\(([0-9]*)\,([0-9]*)\)`)
	matches := mulExpression.FindAllStringSubmatch(corruptedMemory, -1)

	// covert values from string to int and discard whole match, keeping only the selector groups
	var parsedMatches [][]int
	for _, match := range matches {
		var parsedMatch []int
		for _, arg := range match[1:] {
			parsedarg, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			parsedMatch = append(parsedMatch, parsedarg)
		}
		parsedMatches = append(parsedMatches, parsedMatch)
	}

	// multiply and add results
	result := 0
	for _, parsedMatch := range parsedMatches {
		mul := 1
		for _, arg := range parsedMatch {
			mul *= arg
		}
		result += mul
	}

	fmt.Println(result)
}
