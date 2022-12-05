package main

import (
	"strings"
	"testing"
)

const input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	r := strings.NewReader(input)
	result := Part1(r)

	if result != 157 {
		t.Errorf("Expected 157, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input)
	result := Part2(r)

	if result != 157 {
		t.Errorf("Expected 157, got %d", result)
	}
}
