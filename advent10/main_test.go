package main

import (
	"testing"
)

// TestExample1Part1 comprova el exemple més bàsic
func TestExample1Part1(t *testing.T) {

	mapa1 := [][]string{
		{".", "#", ".", ".", "#"},
		{".", ".", ".", ".", "."},
		{"#", "#", "#", "#", "#"},
		{".", ".", ".", ".", "#"},
		{".", ".", ".", "#", "#"},
	}

	expectedBest := Coord{3, 4}

	resultat, best := part1(mapa1)

	if resultat != 8 {
		t.Errorf("Exemple1 dóna %d; want 8", resultat)
	}

	if best != expectedBest {
		t.Errorf("Exemple1 dóna (%v) i hauria de donar (%v)", best, expectedBest)
	}
}

func TestExample2Part1(t *testing.T) {

	linies := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	expectedBest := Coord{5, 8}

	mapa1 := generaMapa(linies)

	resultat, best := part1(mapa1)

	if resultat != 32 {
		t.Errorf("Exemple1 dóna %d; want 32", resultat)
	}

	if best != expectedBest {
		t.Errorf("Exemple1 dóna (%v) i hauria de donar (%v)", best, expectedBest)
	}
}

func TestExample3Part1(t *testing.T) {

	linies := []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###.",
	}
	expectedBest := Coord{1, 2}

	mapa1 := generaMapa(linies)

	resultat, best := part1(mapa1)

	if resultat != 35 {
		t.Errorf("Exemple1 dóna %d; want 35", resultat)
	}
	if best != expectedBest {
		t.Errorf("Exemple1 dóna (%v) i hauria de donar (%v)", best, expectedBest)
	}
}

func TestExample4Part1(t *testing.T) {

	linies := []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#..",
	}
	expectedBest := Coord{6, 3}

	mapa1 := generaMapa(linies)

	resultat, best := part1(mapa1)

	if resultat != 41 {
		t.Errorf("Exemple1 dóna %d; want 41", resultat)
	}

	if best != expectedBest {
		t.Errorf("Exemple1 dóna (%v) i hauria de donar (%v)", best, expectedBest)
	}
}

func TestExample5Part1(t *testing.T) {

	linies := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}
	expectedBest := Coord{11, 13}

	mapa1 := generaMapa(linies)

	resultat, best := part1(mapa1)

	if resultat != 210 {
		t.Errorf("Exemple1 dóna %d; want 210", resultat)
	}

	if best != expectedBest {
		t.Errorf("Exemple1 dóna (%v) i hauria de donar (%v)", best, expectedBest)
	}
}

func TestExample1Part2(t *testing.T) {

	linies := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}

	mapa1 := generaMapa(linies)
	origin := Coord{11, 13}

	resultat := part2(origin, mapa1)

	if resultat != 802 {
		t.Errorf("Exemple1 dóna %d; want 802", resultat)
	}
}
