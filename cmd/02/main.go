package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	sum := Part1(f)
	fmt.Printf("Part 1 solution: %d", sum)
}

func Part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	sum := 0

	for s.Scan() {
		round := strings.Split(s.Text(), " ")
		o := ParseOpponent(round[0])
		self := ParseSelf(round[1])
		sum += self.Fight(o)
	}

	return sum
}

func Part2() int {
	return 1
}

func ParseOpponent(s string) Shape {
	switch s {
	case "A":
		return Rock{}
	case "B":
		return Paper{}
	case "C":
		return Scissors{}
	default:
		log.Fatal("Error parsing opponent")
		return nil
	}
}

func ParseSelf(s string) Shape {
	switch s {
	case "X":
		return Rock{}
	case "Y":
		return Paper{}
	case "Z":
		return Scissors{}
	default:
		log.Fatal("Error parsing opponent")
		return nil
	}
}

type Shape interface {
	Fight(s Shape) int
}

type Rock struct{}

func (r Rock) Fight(s Shape) int {
	total := 1

	switch s.(type) {
	case Rock:
		total += 3

	case Paper:
		total += 0

	case Scissors:
		total += 6

	default:
		log.Fatal("Error Fighting")
	}

	return total
}

type Paper struct{}

func (r Paper) Fight(s Shape) int {
	total := 2

	switch s.(type) {
	case Rock:
		total += 6

	case Paper:
		total += 3

	case Scissors:
		total += 0

	default:
		log.Fatal("Error Fighting")
	}

	return total
}

type Scissors struct{}

func (r Scissors) Fight(s Shape) int {
	total := 3

	switch s.(type) {
	case Rock:
		total += 0

	case Paper:
		total += 6

	case Scissors:
		total += 3

	default:
		log.Fatal("Error Fighting")
	}

	return total

}
