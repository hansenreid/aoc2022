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

	sum := Part2(f)
	fmt.Printf("Part 2 solution: %d", sum)
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

func Part2(r io.Reader) int {
	s := bufio.NewScanner(r)
	sum := 0

	for s.Scan() {
		round := strings.Split(s.Text(), " ")
		o := ParseOpponent(round[0])
		result := ParseResult(round[1])
		sum += result.Rig(o)
	}

	return sum
}

type Result interface {
	Rig(s Shape) int
}

type Win struct{}
type Lose struct{}
type Draw struct{}

func (w Win) Rig(s Shape) int {
	total := 6
	switch s.(type) {
	case Rock:
		return addBase(total, Paper{})
	case Paper:
		return addBase(total, Scissors{})
	case Scissors:
		return addBase(total, Rock{})
	default:
		log.Fatal("Error Rigging Fight")
		return total
	}
}

func (l Lose) Rig(s Shape) int {
	total := 0
	switch s.(type) {
	case Rock:
		return addBase(total, Scissors{})
	case Paper:
		return addBase(total, Rock{})
	case Scissors:
		return addBase(total, Paper{})
	default:
		log.Fatal("Error Rigging Fight")
		return total
	}
}

func (d Draw) Rig(s Shape) int {
	total := 3
	switch s.(type) {
	case Rock:
		return addBase(total, Rock{})
	case Paper:
		return addBase(total, Paper{})
	case Scissors:
		return addBase(total, Scissors{})
	default:
		log.Fatal("Error Rigging Fight")
		return total
	}
}

func addBase(n int, s Shape) int {
	switch s.(type) {
	case Rock:
		n += 1
	case Paper:
		n += 2
	case Scissors:
		n += 3
	default:
		log.Fatal("Invalid Shape")
	}

	return n
}

func ParseResult(s string) Result {
	switch s {
	case "X":
		return Lose{}
	case "Y":
		return Draw{}
	case "Z":
		return Win{}
	default:
		log.Fatal("Error parsing result")
		return nil
	}
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

func (p Paper) Fight(s Shape) int {
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

func (sc Scissors) Fight(s Shape) int {
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
