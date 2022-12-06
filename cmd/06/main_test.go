package main

import (
	"testing"
)

var input = map[string]int{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
	"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
	"nppdvjthqldpwncqszvftbrmjlhg":      6,
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
}

func TestPart1(t *testing.T) {
	for signal, expected := range input {
		result := Part1(signal)
		if result != expected {
			t.Errorf("Expected: %d, Got: %d", expected, result)
		}
	}
}
