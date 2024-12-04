package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		value1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, value1)
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		list2 = append(list2, value2)
	}

	similarityScore := 0
	for _, v := range list1 {
		similarityScore += CountOccurrences(list2, v) * v
	}

	fmt.Println(similarityScore)
}

func CountOccurrences(arr []int, val int) int {
	count := 0
	for _, v := range arr {
		if v == val {
			count++
		}
	}
	return count
}
