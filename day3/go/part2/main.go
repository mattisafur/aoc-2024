package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type parsedMatch struct {
	startIndex int
	args       []int
}

func main() {
	inputBytes, err := os.ReadFile("../../input.txt")
	if err != nil {
		panic(err)
	}
	corruptedMemory := string(inputBytes)

	mulExp := regexp.MustCompile(`mul\(([0-9]*)\,([0-9]*)\)`)
	doExp := regexp.MustCompile(`do\(\)`)
	dontExp := regexp.MustCompile(`don't\(\)`)

	matches := mulExp.FindAllStringSubmatchIndex(corruptedMemory, -1)

	var parsedMatches []parsedMatch
	for _, match := range matches {
		var currentParsedMatch parsedMatch
		currentParsedMatch.startIndex = match[0]
		for i := 2; i < len(match); i += 2 {
			arg := corruptedMemory[match[i]:match[i+1]]
			parsedArg, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			currentParsedMatch.args = append(currentParsedMatch.args, parsedArg)
		}
		parsedMatches = append(parsedMatches, currentParsedMatch)
	}

	var result int
	for _, currentParsedMatch := range parsedMatches {
		corruptedMemoryUpToMul := corruptedMemory[:currentParsedMatch.startIndex]

		doCommandMatches := doExp.FindAllStringIndex(corruptedMemoryUpToMul, -1)
		dontCommandMatches := dontExp.FindAllStringIndex(corruptedMemoryUpToMul, -1)

		if dontCommandMatches != nil && (doCommandMatches == nil || dontCommandMatches[len(dontCommandMatches)-1][1] > doCommandMatches[len(doCommandMatches)-1][1]) {
			continue
		}

		mul := 1
		for _, arg := range currentParsedMatch.args {
			mul *= arg
		}
		result += mul
	}

	fmt.Println(result)
}
