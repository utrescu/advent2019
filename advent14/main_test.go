package main

import (
	"testing"
)

var part1Tests = []struct {
	filename string
	expected int // expected result
}{
	{"input_test1", 31},
	{"input_test2", 165},
	{"input_test3", 13312},
	{"input_test4", 180697},
	{"input_test5", 2210736},
}

// TestExample1Part1 comprova el exemple més bàsic
func TestExample1Part1(t *testing.T) {
	for _, testCasesPart1 := range part1Tests {
		expectedOres := testCasesPart1.expected
		reactions := getReactions(testCasesPart1.filename)

		ores := part1(reactions, "FUEL", 1)

		if ores != expectedOres {
			t.Errorf("Exemple1 dóna %d; want %d", ores, expectedOres)
		}
	}

}

var part2Tests = []struct {
	filename string
	expected int // expected result
}{
	{"input_test3", 82892753},
	{"input_test4", 5586022},
	{"input_test5", 460664},
}

func TestExample1Part2(t *testing.T) {
	for _, testCasesPart2 := range part2Tests {
		expectedFuel := testCasesPart2.expected
		reactions := getReactions(testCasesPart2.filename)

		fuel := part2(reactions, 1000000000000)

		if fuel != expectedFuel {
			t.Errorf("Exemple1 dóna %d; want %d", fuel, expectedFuel)
		}
	}

}
