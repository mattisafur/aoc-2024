package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var wordSearch [][]rune
	for scanner.Scan() {
		wordSearch = append(wordSearch, []rune(strings.Trim(scanner.Text(), "\n")))
	}

	occurences := 0
	for iv := 1; iv < len(wordSearch)-1; iv++ {
	elementLoop:
		for ih := 1; ih < len(wordSearch)-1; ih++ {
			if wordSearch[iv][ih] != rune('A') {
				continue
			}

			for dv := -1; dv <= 1; dv += 2 {
				for dh := -1; dh <= 1; dh += 2 {
					if wordSearch[iv+dv][ih+dh] != 'M' && wordSearch[iv+dv][ih+dh] != 'S' {
						continue elementLoop
					}
				}
			}

			if wordSearch[iv-1][ih-1] == wordSearch[iv+1][ih+1] {
				continue elementLoop
			}
			if wordSearch[iv-1][ih+1] == wordSearch[iv+1][ih-1] {
				continue elementLoop
			}

			occurences++
		}
	}

	fmt.Println(occurences)
}
