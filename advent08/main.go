package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readIntFileCharByChar("input")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("PART1")
	fmt.Println("----------------")
	fmt.Printf("LAYER = %d\n", part1(25, 6, data))
	fmt.Println()

	fmt.Println("PART2")
	fmt.Println("----------------")
	part2(25, 6, data)
}

func part1(wide int, tall int, data []int) int {
	layerSize := wide * tall

	var layers [][]int

	for i := 0; i < len(data); i += layerSize {
		end := i + layerSize
		if end > len(data) {
			end = len(data)
		}

		layers = append(layers, data[i:end])
	}

	minZeros := layerSize + 1
	bestLayer := -1

	for index, layer := range layers {
		zeros := count(0, layer)
		if zeros < minZeros {
			bestLayer = index
			minZeros = zeros
		}
	}

	ones := count(1, layers[bestLayer])
	twos := count(2, layers[bestLayer])

	return ones * twos
}

func count(candidate int, numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number == candidate {
			sum++
		}
	}
	return sum
}

func readIntFileCharByChar(path string) ([]int, error) {
	var result []int

	filebuffer, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		if data.Text() != "\n" {
			i, err := strconv.Atoi(data.Text())
			if err != nil {
				return nil, err
			}
			result = append(result, i)
		}
	}
	return result, nil
}

func part2(wide int, tall int, data []int) {
	layerSize := wide * tall
	numLayers := len(data) / layerSize

	var pixels [][]int

	for pixelPos := 0; pixelPos < layerSize; pixelPos++ {
		var pixel []int
		primer := pixelPos
		for i := 0; i < numLayers; i++ {
			pixel = append(pixel, data[primer])
			primer = primer + layerSize
		}
		pixels = append(pixels, pixel)
	}

	// Generate result
	var result []string
	for _, values := range pixels {
		for _, value := range values {
			if value != 2 {
				if value == 0 {
					result = append(result, " ")
				} else {
					result = append(result, "X")
				}
				break
			}
		}
	}

	// Paint result
	for index, value := range result {
		if index%wide == 0 {
			fmt.Println()
		}
		fmt.Print(value)
	}
	fmt.Println()

}
