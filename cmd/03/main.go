package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	sum := 0
	rucks := ParseRucksacks(r)

	for i := 0; i < len(rucks); i++ {
		dup := rucks[i].duplicate
		sum += items[dup]
	}

	return sum
}

func Part2(r io.Reader) int {
	return 1
}

func ParseRucksacks(r io.Reader) []Rucksack {
	rucksacks := []Rucksack{}
	s := bufio.NewScanner(r)

	for s.Scan() {
		ruck := Rucksack{
			ItemSet: map[rune]struct{}{},
		}

		items := []rune(s.Text())
		l := len(items) / 2

		ruck.Container1 = items[:l]
		ruck.Container2 = items[l:]

		ruck.FindDuplicate()
		rucksacks = append(rucksacks, ruck)
	}

	return rucksacks
}

func dbg(c1 []rune, c2 []rune) {
	fmt.Println("C1:")
	for i := range c1 {
		fmt.Printf("%c\n", c1[i])

	}

	fmt.Println("C2:")
	for i := range c2 {
		fmt.Printf("%c\n", c2[i])

	}
}

type Rucksack struct {
	Container1 []rune
	Container2 []rune
	ItemSet    map[rune]struct{}
	duplicate  rune
}

func (r *Rucksack) FindDuplicate() {
	for i := 0; i < len(r.Container1); i++ {
		item := r.Container1[i]
		r.ItemSet[item] = struct{}{}
	}

	for i := 0; i < len(r.Container2); i++ {
		item := r.Container2[i]

		if _, ok := r.ItemSet[item]; ok {
			r.duplicate = item
		}
	}
}

var items = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}
