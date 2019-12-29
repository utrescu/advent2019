package main

import (
	"testing"
)

var part1Tests = []struct {
	signal   string
	steps    int
	expected string // expected result
}{
	{"12345678", 1, "48226158"},
	{"12345678", 3, "03415518"},
	{"80871224585914546619083218645595", 100, "24176176"},
	{"19617804207202209144916044189917", 100, "73745418"},
	{"69317163492948606335995924319873", 100, "52432133"},
}

// TestExample1Part1 comprova el exemple més bàsic
func TestExample1Part1(t *testing.T) {
	for _, testCasesPart1 := range part1Tests {
		expectedSignal := testCasesPart1.expected

		signal := part1(testCasesPart1.signal, testCasesPart1.steps)

		if signal != expectedSignal {
			t.Errorf("Exemple1 dóna %s; want %s", signal, expectedSignal)
		}
	}

}

var part2Tests = []struct {
	signal   string
	steps    int
	repeats  int
	expected string // expected result
}{
	{"03036732577212944063491565474664", 100, 10000, "84462026"},
	{"02935109699940807407585447034323", 100, 10000, "78725270"},
	{"03081770884921959731165446850517", 100, 10000, "53553731"},
}

// TestExample1Part1 comprova el exemple més bàsic
func TestExample1Part2(t *testing.T) {
	for _, testCasesPart2 := range part2Tests {
		expectedSignal := testCasesPart2.expected

		signal := part2(testCasesPart2.signal, testCasesPart2.steps, testCasesPart2.repeats)

		if signal != expectedSignal {
			t.Errorf("Exemple Part2 dóna %s; want %s", signal, expectedSignal)
		}
	}

}
