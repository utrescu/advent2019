package main

import (
	"testing"
)

// TestExample1Part1 comprova el exemple més bàsic
func TestExample1Part1(t *testing.T) {

	positions := []*Coord{
		&Coord{-1, 0, 2, 0, 0, 0},
		&Coord{2, -10, -7, 0, 0, 0},
		&Coord{4, -8, 8, 0, 0, 0},
		&Coord{3, 5, -1, 0, 0, 0},
	}
	steps := 10
	expectedEnergy := 179

	energy := part1(steps, positions)

	if energy != expectedEnergy {
		t.Errorf("Exemple1 dóna %d; want %d", energy, expectedEnergy)
	}

}

func TestExample2Part1(t *testing.T) {

	positions := []*Coord{
		&Coord{-8, -10, 0, 0, 0, 0},
		&Coord{5, 5, 10, 0, 0, 0},
		&Coord{2, -7, 3, 0, 0, 0},
		&Coord{9, -8, -3, 0, 0, 0},
	}
	steps := 100
	expectedEnergy := 1940

	energy := part1(steps, positions)

	if energy != expectedEnergy {
		t.Errorf("Exemple1 dóna %d; want %d", energy, expectedEnergy)
	}

}
