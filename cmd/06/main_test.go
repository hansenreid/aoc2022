package main

import (
	"testing"
)

var input1 = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
	"nppdvjthqldpwncqszvftbrmjlhg":      6,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
}

var input2 = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
	"nppdvjthqldpwncqszvftbrmjlhg":      23,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
}

func TestPart1(t *testing.T) {
	for signal, expected := range input1 {
		result := Part1(signal)
		if result != expected {
			t.Errorf("Expected: %d, Got: %d", expected, result)
		}
	}
}

func TestPart2(t *testing.T) {
	for signal, expected := range input2 {
		result := Part2(signal)
		if result != expected {
			t.Errorf("Expected: %d, Got: %d", expected, result)
		}
	}
}
