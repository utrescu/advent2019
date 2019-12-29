package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	initial, err := readFile("input")
	if err != nil {
		panic("Error reading the file")
	}

	fmt.Println("------ PART 1 ------- ")
	fmt.Println("FFT: ", part1(initial, 100))

	fmt.Println("------ PART 2 ------- ")
	fmt.Println("FFT: ", part2(initial, 100, 10000))

}

func readFile(path string) (string, error) {
	value, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	signal := strings.TrimSpace(string(value))
	return signal, nil
}

func part1(signal string, fases int) string {
	pattern := []int{0, 1, 0, -1}

	// Generate phases
	for fase := 0; fase < fases; fase++ {
		newSignal := ""
		// One phase
		for digit := range signal {
			sum := 0
			for index, value := range signal {
				// El primer cop és pattern, el segon pattern de dos en dos,
				// el tercer pattern de tres en tres, ...
				sum += int(value-'0') * pattern[(index+1)/(digit+1)%4]
			}
			// Si el valor és negatiu el mòdul falla
			if sum < 0 {
				sum *= -1
			}
			newSignal += string('0' + sum%10)

		}
		signal = newSignal

	}
	return signal[:8]
}

func part2(signal string, fases int, repeats int) string {
	// Repetir
	signal = strings.Repeat(signal, repeats)

	posicioSolucio, err := strconv.Atoi(signal[:7])
	if err != nil {
		panic("Incorrect signal")
	}

	calculaSumatori := []int{}

	for _, digit := range signal[posicioSolucio:] {
		calculaSumatori = append(calculaSumatori, int(digit-'0'))
	}

	for fase := 0; fase < fases; fase++ {

		sum := 0
		for posicio := len(calculaSumatori) - 1; posicio >= 0; posicio-- {
			sum += calculaSumatori[posicio]
			calculaSumatori[posicio] = sum % 10
		}

	}

	result := []string{}

	for _, number := range calculaSumatori[:8] {
		text := strconv.Itoa(number)
		result = append(result, text)
	}

	return strings.Join(result, "")
}
