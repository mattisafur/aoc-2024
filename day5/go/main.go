package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(inputFilePath string) (rules map[int][]int, updates [][]int) {
	// initialize rules
	rules = make(map[int][]int)

	// read and parse input into strings
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rulesStr := strings.Split(sections[0], "\n")
	updatesStr := strings.Split(sections[1], "\n")

	// parse rules
	for _, ruleStr := range rulesStr {
		var rule []int
		for _, ruleValueStr := range strings.Split(ruleStr, "|") {
			ruleValue, err := strconv.Atoi(ruleValueStr)
			if err != nil {
				panic(err)
			}
			rule = append(rule, ruleValue)
		}
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	// parse updates
	for _, updateStr := range updatesStr {
		var update []int
		for _, pageStr := range strings.Split(updateStr, ",") {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				panic(err)
			}
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func calculatePart1(rules map[int][]int, updates [][]int) int {
	sumOfMids := 0
	for _, update := range updates {
		if isUpdateValid(rules, update) {
			sumOfMids += update[len(update)/2]
		}
	}
	return sumOfMids
}

func isUpdateValid(rules map[int][]int, update []int) bool {
	for i, page := range update {
		if !isPageValid(rules, page, update[:i]) {
			return false
		}
	}
	return true
}

func isPageValid(rules map[int][]int, page int, checkAgainst []int) bool {
	for _, previousPage := range checkAgainst {
		for _, rulePage := range rules[page] {
			if previousPage == rulePage {
				return false
			}
		}
	}
	return true
}

func calculatePart2(rules map[int][]int, updates [][]int) int {
	sumOfMids := 0
	for _, update := range updates {
		if isUpdateValid(rules, update) {
			continue
		}

		i := 1
		for i < len(update) {
			firstInvalidIndex := getFirstInvalidIndex(rules, update[i], update[:i])
			if firstInvalidIndex < 0 {
				i++
				continue
			}
			update[i], update[firstInvalidIndex] = update[firstInvalidIndex], update[i]
			i = firstInvalidIndex
		}

		sumOfMids += update[len(update)/2]
	}
	return sumOfMids
}

func getFirstInvalidIndex(rules map[int][]int, page int, checkAgainst []int) int {
	pageRules := rules[page]
	for i, previousPage := range checkAgainst {
		if slices.Contains(pageRules, previousPage) {
			return i
		}
	}
	return -1
}

func main() {
	rules, updates := parseInput("../input.txt")
	fmt.Println(calculatePart1(rules, updates))
	fmt.Println(calculatePart2(rules, updates))
}
