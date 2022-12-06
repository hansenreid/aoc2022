package main

import (
	"reflect"
	"strings"
	"testing"
)

const input = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input)
	ds, ms := SplitInput(r)

	d := ParseDrawing(ds)
	m := ParseMoves(ms)

	expected := "CMZ"
	result := Part1(d, m)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input)
	ds, ms := SplitInput(r)

	d := ParseDrawing(ds)
	m := ParseMoves(ms)

	expected := "MCD"
	result := Part2(d, m)

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestParseDrawing(t *testing.T) {
	r := strings.NewReader(input)
	s, _ := SplitInput(r)

	expected := Drawing{
		map[int]Stack{
			1: {
				Crates: []Crate{'Z', 'N'},
			},

			2: {
				Crates: []Crate{'M', 'C', 'D'},
			},

			3: {
				Crates: []Crate{'P'},
			},
		},
	}

	d := ParseDrawing(s)

	if !reflect.DeepEqual(d, expected) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, d)
	}
}

func TestParseMove(t *testing.T) {
	r := strings.NewReader(input)
	_, s := SplitInput(r)

	expected := []Move{
		{count: 1, from: 2, to: 1},
		{count: 3, from: 1, to: 3},
		{count: 2, from: 2, to: 1},
		{count: 1, from: 1, to: 2},
	}

	m := ParseMoves(s)

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, m)
	}

}

func TestPutMany(t *testing.T) {
	stack := NewStack()

	expected := Stack{
		Crates: []Crate{'A', 'B', 'C'},
	}

	stack.PutMany([]Crate{'A', 'B', 'C'})

	if !reflect.DeepEqual(stack, expected) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, stack)
	}
}

func TestPopMany(t *testing.T) {
	stack := NewStack()
	stack.PutMany([]Crate{'A', 'B', 'C'})

	stack2 := NewStack()
	stack2.PutMany([]Crate{'Z'})

	expected1 := Stack{
		Crates: []Crate{'A'},
	}

	expected2 := Stack{
		Crates: []Crate{'Z', 'B', 'C'},
	}

	items := stack.PopMany(2)
	stack2.PutMany(items)

	if !reflect.DeepEqual(stack, expected1) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected1, stack)
	}

	if !reflect.DeepEqual(stack2, expected2) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected2, stack2)
	}

}
