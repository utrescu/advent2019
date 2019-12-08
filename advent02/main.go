package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	values, err := readFileIntoArray("input")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("---- PART 1: Exemples ------")
	fmt.Printf("1,0,0,0,99 és %d\n", part1(0, 0, []int{1, 0, 0, 0, 99}))
	fmt.Printf("2,3,0,3,99 és %d\n", part1(3, 0, []int{2, 3, 0, 3, 99}))
	fmt.Printf("2,4,4,5,99,0 és %d\n", part1(4, 4, []int{2, 4, 4, 5, 99, 0}))
	fmt.Printf("1,1,1,4,99,5,6,0,99 és %d\n", part1(1, 1, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
	fmt.Println("-------- PART 1 ------------")
	list := append([]int(nil), values...)
	fmt.Printf("TOTAL: %d\n", part1(12, 2, list))
	fmt.Println("-------- PART 2 ------------")

	for verb := 0; verb <= 99; verb++ {
		for noun := 0; noun <= 99; noun++ {
			list := append([]int(nil), values...)
			result := part1(noun, verb, list)
			if result == 19690720 {
				fmt.Printf("TOTAL: %d\n", noun*100+verb)
				break
			}
		}
	}

}

func readFileIntoArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	line := strings.TrimRight(string(b), "\r\n")
	llistaStrings := strings.Split(line, ",")

	return sliceAtoi(llistaStrings)
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func part1(noun int, verb int, values []int) int {
	opCodePos := 0

	values[1] = noun
	values[2] = verb

	for values[opCodePos] != 99 {

		result := values[opCodePos+3]
		operand1 := values[opCodePos+1]
		operand2 := values[opCodePos+2]

		switch opCode := values[opCodePos]; opCode {
		case 1:
			// Sumar
			values[result] = values[operand1] + values[operand2]
		case 2:
			values[result] = values[operand1] * values[operand2]
		case 99:
		default:
			return -1
		}
		opCodePos += 4
	}
	return values[0]
}
