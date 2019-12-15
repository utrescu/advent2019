package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Coord defines a pos in the map
type Coord struct {
	X int
	Y int
}

// Distance distance onto point^2
func (me Coord) Distance(other Coord) int {
	return (me.X-other.X)*(me.X-other.X) + (me.Y-other.Y)*(me.Y-other.Y)
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
	return lines, scanner.Err()
}

func generaMapa(linies []string) [][]string {
	var resultat [][]string
	for _, linia := range linies {
		var separed []string
		for _, c := range []rune(linia) {
			separed = append(separed, string(c))
		}
		resultat = append(resultat, separed)
	}

	return resultat
}

func main() {

	lines, err := readLines("input")
	if err != nil {
		panic(err)
	}
	fmt.Println("--------- part 1 ----------------")
	mapa := generaMapa(lines)
	result, best := part1(mapa)
	fmt.Printf("Best: %d %v\n", result, best)

	fmt.Println("--------- part 2 ----------------")
	fmt.Printf("Best: %d\n", part2(best, mapa))

}

func getAsteroids(mapa [][]string) []Coord {

	var result []Coord

	for y, line := range mapa {
		for x, value := range line {
			if value == "#" {
				asteroid := Coord{x, y}
				result = append(result, asteroid)
			}
		}
	}
	return result
}

func getVisionFrom(asteroid Coord, asteroids []Coord) map[float64][]Coord {
	angles := make(map[float64][]Coord)
	for _, asteroid2 := range asteroids {
		if asteroid == asteroid2 {
			continue
		}

		angle := -math.Atan2(float64(asteroid2.X-asteroid.X), float64(asteroid2.Y-asteroid.Y))
		if value, ok := angles[angle]; !ok {
			angles[angle] = append(value, asteroid2)
		}
	}
	return angles
}

func part1(mapa [][]string) (int, Coord) {
	asteroids := getAsteroids(mapa)
	max := 0
	best := Coord{-1, -1}
	for _, asteroid := range asteroids {

		angles := getVisionFrom(asteroid, asteroids)
		found := len(angles)
		// fmt.Println(asteroid, "dóna", found)
		if found > max {
			max = found
			best = asteroid
		}
	}
	return max, best
}

func part2(origin Coord, mapa [][]string) int {
	asteroids := getAsteroids(mapa)
	angles := getVisionFrom(origin, asteroids)

	var anglesList []float64
	i := 0
	for i < 200 {

		for k := range angles {
			anglesList = append(anglesList, k)
		}
		sort.Float64s(anglesList)

		for _, angleItem := range anglesList {

			if values, ok := angles[angleItem]; ok {
				i++
				// Obtenir el més proper
				index, candidate := getCloser(origin, values)
				if i == 200 {
					return candidate.X*100 + candidate.Y
				}

				if len(values) == 1 {
					delete(angles, angleItem)
				} else {
					// Eliminar el més proper
					values[index] = values[len(values)-1]
					values = values[:len(values)-1]
					angles[angleItem] = values
				}

			}

		}

	}

	return 0
}

func getCloser(origin Coord, points []Coord) (int, Coord) {
	pos := 0
	best := points[0]
	distance := origin.Distance(best)

	for index, element := range points {
		newDistance := origin.Distance(element)
		if newDistance < distance {
			pos = index
			best = element
			distance = newDistance
		}
	}

	return pos, best
}
