package main

import (
	"strings"
	"testing"
)

const input = `A Y
B X
C Z`

func TestPart(t *testing.T) {
	r := strings.NewReader(input)
	result := Part1(r)

	if result != 15 {
		t.Errorf("Expected 15, got: %d", result)
	}
}

func TestRockVRock(t *testing.T) {
	self := Rock{}
	opponent := Rock{}

	result := self.Fight(opponent)

	if result != 4 {
		t.Errorf("Expect 4, got: %d", result)
	}

}

func TestRockVPaper(t *testing.T) {
	self := Rock{}
	opponent := Paper{}

	result := self.Fight(opponent)

	if result != 1 {
		t.Errorf("Expect 1, got: %d", result)
	}

}

func TestRockVScissors(t *testing.T) {
	self := Rock{}
	opponent := Scissors{}

	result := self.Fight(opponent)

	if result != 7 {
		t.Errorf("Expect 7, got: %d", result)
	}

}

func TestPaperVRock(t *testing.T) {
	self := Paper{}
	opponent := Rock{}

	result := self.Fight(opponent)

	if result != 8 {
		t.Errorf("Expect 8, got: %d", result)
	}

}

func TestPaperVPaper(t *testing.T) {
	self := Paper{}
	opponent := Paper{}

	result := self.Fight(opponent)

	if result != 5 {
		t.Errorf("Expect 5, got: %d", result)
	}

}

func TestPaperVScissors(t *testing.T) {
	self := Paper{}
	opponent := Scissors{}

	result := self.Fight(opponent)

	if result != 2 {
		t.Errorf("Expect 2, got: %d", result)
	}

}
func TestScissorsVRock(t *testing.T) {
	self := Scissors{}
	opponent := Rock{}

	result := self.Fight(opponent)

	if result != 3 {
		t.Errorf("Expect 3, got: %d", result)
	}

}

func TestScissorsVPaper(t *testing.T) {
	self := Scissors{}
	opponent := Paper{}

	result := self.Fight(opponent)

	if result != 9 {
		t.Errorf("Expect 9, got: %d", result)
	}

}

func TestSciccorsVScissors(t *testing.T) {
	self := Scissors{}
	opponent := Scissors{}

	result := self.Fight(opponent)

	if result != 6 {
		t.Errorf("Expect 6, got: %d", result)
	}

}
