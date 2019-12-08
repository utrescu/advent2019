package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	masses, err := readIntLines("input1")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("---- PART 1: Exemples ------")
	fmt.Printf("12 és %d\n", 12/3-2)
	fmt.Printf("14 és %d\n", 14/3-2)
	fmt.Printf("1969 és %d\n", 1969/3-2)
	fmt.Printf("100756 és %d\n", 100756/3-2)
	fmt.Println("---------- PART 1 ----------")
	part1(masses)

	fmt.Println("---- PART 2: Exemples ------")
	fmt.Printf("14 és %d\n", part2([]int{14}))
	fmt.Printf("1969 és %d\n", part2([]int{1969}))
	fmt.Printf("100756 és %d\n", part2([]int{100756}))
	fmt.Println("---------- PART 2 ----------")
	fmt.Printf("TOTAL és %d\n", part2(masses))
}

func readIntLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(fmt.Sprintf("No es pot convertir %s", scanner.Text()))
		}
		lines = append(lines, newMass)
	}
	return lines, scanner.Err()
}

func part1(masses []int) {

	mass := 0

	for _, newMass := range masses {
		mass += newMass/3 - 2
	}

	fmt.Printf("Total: %d\n", mass)
}

func part2(masses []int) int {

	fuel := 0

	for _, newMass := range masses {
		fuel += calculaFuel(newMass)
	}

	return fuel
}

func calculaFuel(fuel int) int {
	mass := 0

	for fuel > 0 {
		newMass := fuel/3 - 2
		fuel = newMass
		if newMass > 0 {
			mass += fuel
		}
	}
	return mass
}
