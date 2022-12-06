package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	s := bufio.NewScanner(f)
	s.Scan()
	input := s.Text()

	fmt.Println("Part 1 solution: ", Part1(input))

}

func Part1(s string) int {
	return FirstUnique(s)
}

func FirstUnique(s string) int {
	runes := []rune(s)
	var window []rune
	var windowSet map[rune]struct{}

	for i := 0; i < len(runes)-3; i++ {
		window = runes[i : i+4]
		windowSet = make(map[rune]struct{}, 4)
		for _, v := range window {
			windowSet[v] = struct{}{}
		}

		if len(windowSet) == 4 {
			return i + 4
		}
	}

	return 0
}
