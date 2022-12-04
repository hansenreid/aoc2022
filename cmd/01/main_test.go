package main

import (
	"reflect"
	"strings"
	"testing"
)

var input string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input)
	result := part1(r)

	if result != 24000 {
		t.Errorf("Expected 24000, got %d", result)
	}
}

func TestNewElf(t *testing.T) {
	input := []string{"1000", "2000", "3000"}

	elf := NewElf(input)

	expected := &Elf{
		Calories: []int{1000, 2000, 3000},
		Total:    6000,
	}

	if !reflect.DeepEqual(elf, expected) {
		t.Errorf("Expected %v, got %v", expected, elf)
	}
}

func TestParseElves(t *testing.T) {
	r := strings.NewReader(input)
	elves := ParseElves(r)

	expected := []Elf{
		{
			Calories: []int{1000, 2000, 3000},
			Total:    6000,
		},
		{
			Calories: []int{4000},
			Total:    4000,
		},
		{
			Calories: []int{5000, 6000},
			Total:    11000,
		},
		{
			Calories: []int{7000, 8000, 9000},
			Total:    24000,
		},
		{
			Calories: []int{10000},
			Total:    10000,
		},
	}

	if !reflect.DeepEqual(elves, expected) {
		t.Errorf("Expected %v, got %v", expected, elves)
	}
}
