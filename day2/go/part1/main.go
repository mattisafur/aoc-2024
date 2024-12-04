package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read and parse input
	file, err := os.Open("../../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels_str := strings.Fields(scanner.Text())

		var report []int
		for _, level_str := range levels_str {
			level, err := strconv.Atoi(level_str)
			if err != nil {
				panic(err)
			}
			report = append(report, level)
		}
		data = append(data, report)
	}

	safeReportsCount := 0
	for _, report := range data {
		if isValid(report) {
			safeReportsCount++
		}
	}

	fmt.Println(safeReportsCount)
}

func isValid(report []int) bool {
	reportAscendingValid := true
	reportDescendingValid := true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 1 || diff > 3 {
			reportAscendingValid = false
		}
		if diff < -3 || diff > -1 {
			reportDescendingValid = false
		}
	}
	return reportAscendingValid || reportDescendingValid
}
