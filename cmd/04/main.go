package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	a := ParseAssignments(f)

	sum := Part1(a)
	fmt.Printf("Part 1 solution: %d", sum)

}

func Part1(a []Assignment) int {
	sum := 0
	for i := range a {
		elf1 := a[i].elf1
		elf2 := a[i].elf2

		if contains(&elf1, &elf2) {
			sum += 1
		}
	}

	return sum
}

func ParseAssignments(r io.Reader) []Assignment {
	assignments := []Assignment{}
	s := bufio.NewScanner(r)

	for s.Scan() {
		input := strings.Split(s.Text(), ",")
		elf1 := ParseRange(input[0])
		elf2 := ParseRange(input[1])

		a := Assignment{
			elf1: elf1,
			elf2: elf2,
		}

		assignments = append(assignments, a)
	}

	return assignments
}

func ParseRange(s string) Elf {
	input := strings.Split(s, "-")

	start, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal("Error parsing range")
	}

	end, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal("Error parsing range")
	}

	return Elf{
		start: start,
		end:   end,
	}
}

type Assignment struct {
	elf1 Elf
	elf2 Elf
}

type Elf struct {
	start int
	end   int
}

func contains(elf1 *Elf, elf2 *Elf) bool {
	if elf1.start <= elf2.start && elf1.end >= elf2.end {
		return true
	}

	if elf2.start <= elf1.start && elf2.end >= elf1.end {
		return true
	}

	return false
}
