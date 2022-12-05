package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	elves := ParseElves(f)

	fmt.Printf("Part 1 solution: %d\n", part1(elves))
	fmt.Printf("Part 2 solution: %d\n", part2(elves))
}

func part1(elves []Elf) int {
	max := 0

	for i := 0; i < len(elves); i++ {
		if elves[i].Total > max {
			max = elves[i].Total
		}
	}

	return max
}

func part2(elves []Elf) int {
	sort.Sort(ByTotal(elves))

	total := 0
	for i := range elves[:3] {
		total += elves[i].Total
	}

	return total
}

type Elf struct {
	Calories []int
	Total    int
}

type ByTotal []Elf

func (e ByTotal) Len() int           { return len(e) }
func (e ByTotal) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e ByTotal) Less(i, j int) bool { return e[j].Total < e[i].Total }

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
