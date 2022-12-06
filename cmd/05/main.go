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

	ds, ms := SplitInput(f)

	d := ParseDrawing(ds)
	m := ParseMoves(ms)

	result := Part2(d, m)

	fmt.Println("Part 2 result: ", result)

}

func Part1(d Drawing, moves []Move) string {
	for i := range moves {
		d.Move(moves[i])
	}

	top := ""

	for i := 0; i < len(d.Stacks); i++ {
		top += string(d.Stacks[i+1].Peek())
	}

	return top
}

func Part2(d Drawing, moves []Move) string {
	for i := range moves {
		fmt.Println("Moving:\n", moves[i], "\n", d)
		d.MoveMany(moves[i])
	}

	top := ""

	for i := 0; i < len(d.Stacks); i++ {
		top += string(d.Stacks[i+1].Peek())
	}

	return top
}

func SplitInput(r io.Reader) (drawing, commands []string) {
	d := []string{}
	c := []string{}

	finishedDrawing := false

	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()

		if text == "" {
			finishedDrawing = true
			continue
		}

		if !finishedDrawing {
			d = append(d, text)
			continue
		}

		c = append(c, text)
	}

	return d, c
}

type Drawing struct {
	Stacks map[int]Stack
}

func NewDrawing() Drawing {
	return Drawing{
		Stacks: map[int]Stack{},
	}
}

func ParseDrawing(s []string) Drawing {
	d := NewDrawing()

	for _, line := range s {
		if strings.HasPrefix(line, " 1") {
			return d
		}

		d.AddCrates(line)
	}

	log.Fatal("Did not find end of drawing")
	return d
}

func (d *Drawing) AddCrates(s string) {
	runes := []rune(s)

	for i, r := range runes {
		if r != ' ' && r != '[' && r != ']' {
			n := (i-1)/4 + 1
			stack := d.Stacks[n]
			stack.Prepend(Crate(r))

			d.Stacks[n] = stack
		}
	}
}

func (d *Drawing) Move(m Move) {
	fromStack := d.Stacks[m.from]
	toStack := d.Stacks[m.to]

	for i := 0; i < m.count; i++ {
		crate := fromStack.Pop()
		toStack.Put(crate)
	}

	d.Stacks[m.from] = fromStack
	d.Stacks[m.to] = toStack
}

func (d *Drawing) MoveMany(m Move) {
	fromStack := d.Stacks[m.from]
	toStack := d.Stacks[m.to]

	crates := fromStack.PopMany(m.count)
	toStack.PutMany(crates)

	d.Stacks[m.from] = fromStack
	d.Stacks[m.to] = toStack
}

type Stack struct {
	Crates []Crate
}

func NewStack() Stack {
	return Stack{
		Crates: []Crate{},
	}
}

func (s *Stack) Put(c Crate) {
	if s == nil {
		*s = NewStack()
	}

	s.Crates = append(s.Crates, c)
}

func (s *Stack) Prepend(c Crate) {
	if s == nil {
		*s = NewStack()
	}

	s.Crates = append([]Crate{c}, s.Crates...)
}

func (s *Stack) Pop() Crate {
	l := len(s.Crates)

	if l == 0 {
		log.Fatal("Attempting to Pop from empty stack")
	}

	item := s.Crates[l-1]
	s.Crates = s.Crates[:l-1]

	return item
}

func (s *Stack) PopMany(n int) []Crate {
	l := len(s.Crates)

	if l < n {
		fmt.Printf("Stack: %v\nCount: %d\n", s, n)
		log.Fatal("Attempting to Pop too many from stack")
	}

	items := s.Crates[l-n:]
	s.Crates = s.Crates[:l-n]

	return items
}

func (s *Stack) PutMany(c []Crate) {
	if s == nil {
		*s = NewStack()
	}

	s.Crates = append(s.Crates, c...)
}

func (s Stack) Peek() Crate {
	l := len(s.Crates)

	if l == 0 {
		log.Fatal("Peeking an empty stack")
	}

	return s.Crates[l-1]
}

type Crate rune

func (c Crate) String() string {
	return fmt.Sprintf("%c", c)
}

type Move struct {
	count int
	from  int
	to    int
}

func ParseMoves(s []string) []Move {
	moves := []Move{}

	for _, line := range s {
		split := strings.Split(line, " ")

		count, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Error parsing count")
		}

		from, err := strconv.Atoi(split[3])
		if err != nil {
			log.Fatal("Error parsing from")
		}

		to, err := strconv.Atoi(split[5])
		if err != nil {
			log.Fatal("Error parsing to")
		}

		moves = append(moves, Move{count: count, from: from, to: to})
	}

	return moves
}
