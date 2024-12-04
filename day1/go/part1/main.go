package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// read values
	file, err := os.Open("../../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1, list2 []int

	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		num1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// sort values
	slices.Sort(list1)
	slices.Sort(list2)

	// calculate sum of diffs
	diff_sum := 0
	for i := 0; i < len(list1) && i < len(list2); i++ {
		if list1[i] > list2[i] {
			diff_sum += list1[i] - list2[i]
		} else {
			diff_sum += list2[i] - list1[i]
		}
	}

	fmt.Println(diff_sum)
}
