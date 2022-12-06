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
	fmt.Println("Part 2 solution: ", Part2(input))

}

func Part1(s string) int {
	return FirstUnique(s, 4)
}

func Part2(s string) int {
	return FirstUnique(s, 14)
}

func FirstUnique(s string, n int) int {
	runes := []rune(s)
	var window []rune
	var windowSet map[rune]struct{}

	for i := 0; i < len(runes)-n+1; i++ {
		window = runes[i : i+n]
		windowSet = make(map[rune]struct{}, n)
		for _, v := range window {
			windowSet[v] = struct{}{}
		}

		if len(windowSet) == n {
			return i + n
		}
	}

	return 0
}
