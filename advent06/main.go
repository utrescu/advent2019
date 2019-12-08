package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Orbits representa el planeta i el satèl·lit
type Orbits struct {
	Orbited string
	Orbiter string
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		panic(err.Error())
	}

	cas1 := []Orbits{
		Orbits{"COM", "B"},
		Orbits{"B", "C"},
		Orbits{"C", "D"},
		Orbits{"D", "E"},
		Orbits{"E", "F"},
		Orbits{"B", "G"},
		Orbits{"G", "H"},
		Orbits{"D", "I"},
		Orbits{"E", "J"},
		Orbits{"J", "K"},
		Orbits{"K", "L"}}

	fmt.Printf("cas1 dóna %d i ha de donar 42\n", part1(cas1))

	fmt.Println("---------- PART 1 ----------")
	fmt.Printf("ORBITS és %d - 308790!\n", part1(lines))

	fmt.Println("---------- Exemple PART 2 ----------")
	cas2 := []Orbits{
		Orbits{"COM", "B"},
		Orbits{"B", "C"},
		Orbits{"C", "D"},
		Orbits{"D", "E"},
		Orbits{"E", "F"},
		Orbits{"B", "G"},
		Orbits{"G", "H"},
		Orbits{"D", "I"},
		Orbits{"E", "J"},
		Orbits{"J", "K"},
		Orbits{"K", "L"},
		Orbits{"K", "YOU"},
		Orbits{"I", "SAN"}}

	fmt.Printf("cas2 dóna %d i ha de donar 4\n", part2("YOU", "SAN", cas2))

	fmt.Println("---------- PART 2 ----------")
	fmt.Printf("JUMPS és %d - ?!\n", part2("YOU", "SAN", lines))
}

func readLines(path string) ([]Orbits, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Orbits
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		planets := strings.Split(line, ")")
		lines = append(lines, Orbits{planets[0], planets[1]})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func part1(orbits []Orbits) uint {

	var candidates []string

	planetOrbits := make(map[string]uint)

	// Crear el mapa de resultats
	for _, orbit := range orbits {

		if _, ok := planetOrbits[orbit.Orbited]; !ok {
			planetOrbits[orbit.Orbited] = 0
		} else {
			planetOrbits[orbit.Orbited] = 1
		}

		if _, ok := planetOrbits[orbit.Orbiter]; !ok {
			planetOrbits[orbit.Orbiter] = 1
		}
	}

	// S'ha de començar pel que no orbiti
	for key, value := range planetOrbits {
		if value == 0 {
			candidates = append(candidates, key)
		}
	}

	// Start with root
	// fmt.Printf("... Root planet %s\n", root)

	// Process orbited -> orbiter until no more candidates

	numberCandidates := len(candidates)

	for numberCandidates > 0 {
		// get candidate
		candidate := candidates[0]
		// fmt.Printf("....... Processing %s\n", candidate)
		candidates = candidates[1:]
		orbiters := findOrbiters(candidate, orbits)
		for _, orbiter := range orbiters {

			planetOrbits[orbiter] = 1 + planetOrbits[candidate]
			candidates = append(candidates, orbiter)
		}
		numberCandidates = len(candidates)
	}

	// Sum orbits
	var total uint = 0
	for _, value := range planetOrbits {
		total = total + value
	}

	return total
}

func findOrbiters(candidate string, orbits []Orbits) []string {
	var result []string
	for _, orbit := range orbits {
		if orbit.Orbited == candidate {
			result = append(result, orbit.Orbiter)
		}
	}
	return result
}

// part2 busca els salts
func part2(you string, san string, orbits []Orbits) int {

	pathToYou := pathTo(you, orbits)
	pathToSan := pathTo(san, orbits)
	result := difference(pathToYou, pathToSan)

	return len(result)
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func pathTo(desti string, orbits []Orbits) []string {
	var solution []string

	next := desti
	for next != "END" {
		destiBefore := next
		for _, value := range orbits {
			if value.Orbiter == next {
				next = value.Orbited
				solution = append(solution, next)
				break
			}
		}
		if next == destiBefore {
			next = "END"
		}
	}

	return solution
}
