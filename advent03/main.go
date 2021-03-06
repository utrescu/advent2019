package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Coord represents cartesian Map coords
type Coord struct {
	X int
	Y int
}

// Move moves a Coord
func (value *Coord) Move(inc Coord) {
	value.X += inc.X
	value.Y += inc.Y
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("---- PART 1: Exemples ------")
	cas0 := []string{
		"R8,U5,L5,D3",
		"U7,R6,D4,L4"}
	one, _ := part1(cas0)
	fmt.Printf("cas0 dóna %d i ha de donar 6\n", one)

	cas1 := []string{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83"}

	one, _ = part1(cas1)
	fmt.Printf("cas1 dóna %d i ha de donar 159\n", one)

	cas2 := []string{
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}

	one, _ = part1(cas2)
	fmt.Printf("cas2 dóna %d i ha de donar 135\n", one)

	fmt.Println("---------- PART 1 ----------")
	one, _ = part1(lines)
	fmt.Printf("DISTANCE és %d\n", one)

	fmt.Println("---- PART 2: Exemples ------")

	_, two := part1(cas0)
	fmt.Printf("cas0 dóna %d i ha de donar 30\n", two)

	_, two = part1(cas1)
	fmt.Printf("cas1 dóna %d i ha de donar 610\n", two)

	_, two = part1(cas2)
	fmt.Printf("cas2 dóna %d i ha de donar 410\n", two)

	fmt.Println("---------- PART 2 ----------")
	_, two = part1(lines)
	fmt.Printf("DISTANCE és %d\n", two)

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(lines []string) (int, int) {

	direction := Coord{0, 0}

	world := make(map[Coord]int)
	creuaments := make(map[Coord]int)

	var best = math.MaxInt32

	for lineNum, line := range lines {
		steps := 0
		initial := &Coord{0, 0}
		valors := strings.Split(line, ",")

		for pos, valor := range valors {
			runes := []rune(valor)
			switch move := string(runes[0]); move {
			case "R":
				direction = Coord{1, 0}
			case "L":
				direction = Coord{-1, 0}
			case "U":
				direction = Coord{0, -1}
			case "D":
				direction = Coord{0, 1}
			default:
				panic(fmt.Sprintf("Unknown direction %s of %d:%s from %s", move, pos, valor, line))
			}

			times, err := strconv.Atoi(string(runes[1:]))
			if err != nil {
				panic(err.Error())
			}
			for i := 0; i < times; i++ {

				initial.Move(direction)
				steps = steps + 1

				if _, ok := world[*initial]; ok {
					// Ja hi era
					if lineNum != 0 {
						// Part 1
						distance := abs(initial.X) + abs(initial.Y)
						if distance < best {
							best = distance
						}
						actual := world[*initial]
						world[*initial] = actual + steps

						creuaments[*initial] = actual + steps
					}
					// Part 2

				} else {
					world[*initial] = steps
				}
			}
		}
	}

	// Part 2
	minCombinedSteps := -1
	for val, combinedSteps := range creuaments {

		if minCombinedSteps < 0 || combinedSteps < minCombinedSteps {
			minCombinedSteps = combinedSteps
			fmt.Printf(" -> %d,%d: %d steps\n", val.X, val.Y, combinedSteps)
		}
	}

	return best, minCombinedSteps
}
