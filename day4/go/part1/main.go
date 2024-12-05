package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read word search
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

	word := []rune("XMAS")

	occurences := 0
	for iv := 0; iv < len(wordSearch); iv++ {
		for ih := 0; ih < len(wordSearch[0]); ih++ {
			possibleVerticalMoves := []int{0}
			possibleHorizontalMoves := []int{0}
			if iv > len(word)-2 {
				possibleVerticalMoves = append(possibleVerticalMoves, -1)
			}
			if iv < len(wordSearch)-len(word)+1 {
				possibleVerticalMoves = append(possibleVerticalMoves, 1)
			}
			if ih > len(word)-2 {
				possibleHorizontalMoves = append(possibleHorizontalMoves, -1)
			}
			if ih < len(wordSearch[0])-len(word)+1 {
				possibleHorizontalMoves = append(possibleHorizontalMoves, 1)
			}

			for _, dv := range possibleVerticalMoves {
				for _, dh := range possibleHorizontalMoves {
					if dv == 0 && dh == 0 {
						continue
					}

					isOccurence := true
					for i := 0; i < len(word); i++ {
						if wordSearch[iv+dv*i][ih+dh*i] != word[i] {
							isOccurence = false
							break
						}
					}
					if isOccurence {
						occurences++
					}
				}
			}
		}
	}

	fmt.Println(occurences)
}
