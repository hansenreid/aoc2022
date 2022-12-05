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

	sum := Part2(f)
	fmt.Printf("Part 2 solution: %d", sum)

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
	rucks := ParseRucksacks(r)
	sum := 0

	for i := 0; i < len(rucks); i += 3 {
		a := rucks[i]
		b := rucks[i+1]
		c := rucks[i+2]

		badge := FindBadge(&a, &b, &c)
		sum += items[badge]
	}

	return sum
}

func ParseRucksacks(r io.Reader) []Rucksack {
	rucksacks := []Rucksack{}
	s := bufio.NewScanner(r)

	for s.Scan() {
		ruck := Rucksack{
			DuplicateSet: map[rune]struct{}{},
			ContainerSet: map[rune]struct{}{},
		}

		items := []rune(s.Text())
		l := len(items) / 2

		ruck.SetContainers(items[:l], items[l:])

		ruck.FindDuplicate()
		rucksacks = append(rucksacks, ruck)
	}

	return rucksacks
}

func FindBadge(a *Rucksack, b *Rucksack, c *Rucksack) rune {
	for k := range a.ContainerSet {
		_, ok1 := b.ContainerSet[k]
		_, ok2 := c.ContainerSet[k]
		if ok1 && ok2 {
			return k
		}
	}

	log.Fatal("Error finding Badge")
	return 0
}

type Rucksack struct {
	Container1   []rune
	Container2   []rune
	ContainerSet map[rune]struct{}
	DuplicateSet map[rune]struct{}
	duplicate    rune
}

func (r *Rucksack) SetContainers(c1 []rune, c2 []rune) {
	r.Container1 = c1
	r.Container2 = c2

	for i := 0; i < len(c1); i++ {
		r.ContainerSet[c1[i]] = struct{}{}
	}

	for i := 0; i < len(c2); i++ {
		r.ContainerSet[c2[i]] = struct{}{}
	}
}

func (r *Rucksack) FindDuplicate() {
	for i := 0; i < len(r.Container1); i++ {
		item := r.Container1[i]
		r.DuplicateSet[item] = struct{}{}
	}

	for i := 0; i < len(r.Container2); i++ {
		item := r.Container2[i]

		if _, ok := r.DuplicateSet[item]; ok {
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
