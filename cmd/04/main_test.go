package main

import (
	"reflect"
	"strings"
	"testing"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input)
	a := ParseAssignments(r)

	result := Part1(a)

	if result != 2 {
		t.Errorf("Expected 2, got: %d", result)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input)
	a := ParseAssignments(r)

	result := Part2(a)

	if result != 4 {
		t.Errorf("Expected 4, got: %d", result)
	}
}

func TestParseAssignments(t *testing.T) {
	r := strings.NewReader(input)
	expected := []Assignment{
		{
			elf1: Elf{start: 2, end: 4},
			elf2: Elf{start: 6, end: 8},
		},
		{
			elf1: Elf{start: 2, end: 3},
			elf2: Elf{start: 4, end: 5},
		},
		{
			elf1: Elf{start: 5, end: 7},
			elf2: Elf{start: 7, end: 9},
		},
		{
			elf1: Elf{start: 2, end: 8},
			elf2: Elf{start: 3, end: 7},
		},
		{
			elf1: Elf{start: 6, end: 6},
			elf2: Elf{start: 4, end: 6},
		},
		{
			elf1: Elf{start: 2, end: 6},
			elf2: Elf{start: 4, end: 8},
		},
	}

	result := ParseAssignments(r)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%v\nGot:\n%v\n", expected, result)
	}

}
