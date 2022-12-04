package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Printf("Part 1 solution: %d", part1(f))
}

func part1(r io.Reader) int {
	max := 0
	elves := ParseElves(r)

	for i := 0; i < len(elves); i++ {

		if elves[i].Total > max {
			max = elves[i].Total
		}

	}

	return max
}

type Elf struct {
	Calories []int
	Total    int
}

func ParseElves(r io.Reader) []Elf {
	s := bufio.NewScanner(r)
	currInput := []string{}
	elves := []Elf{}

	for s.Scan() {
		text := s.Text()
		if text == "" {

			elves = append(elves, *NewElf(currInput))
			currInput = []string{}
			continue
		}

		currInput = append(currInput, text)
	}

	elves = append(elves, *NewElf(currInput))

	return elves
}

func NewElf(input []string) *Elf {
	calories := []int{}

	for i := 0; i < len(input); i++ {

		calorie, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatal("Failed to parse calorie: ", input[i])
		}

		calories = append(calories, calorie)
	}

	sum := 0
	for i := 0; i < len(calories); i++ {
		sum += calories[i]
	}

	elf := &Elf{
		Calories: calories,
		Total:    sum,
	}

	return elf
}
