package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Reaction represents the quimical reaction
type Reaction struct {
	Quantity    int
	Ingredients map[string]int
}

func main() {
	fmt.Println("------ PART 1 ------- ")
	reactions := getReactions("input")
	fmt.Println("ORES: ", part1(reactions, "FUEL", 1))

	fmt.Println("------ PART 2 ------- ")
	reactions = getReactions("input")
	trillion := 1000000000000
	fmt.Println("FUEL: ", part2(reactions, trillion))

}

func getReactions(filename string) map[string]Reaction {

	input, _ := ioutil.ReadFile(filename)

	reactions := map[string]Reaction{}

	for _, value := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		reactionString := strings.Split(value, " => ")
		result := strings.Split(reactionString[1], " ")

		// Element result
		ingredientName := result[1]
		ingredientQuantity, _ := strconv.Atoi(result[0])
		reactions[ingredientName] = Reaction{ingredientQuantity, map[string]int{}}

		// Ingredients
		for _, s := range strings.Split(reactionString[0], ", ") {
			inputIngredient := strings.Split(s, " ")
			quantity, _ := strconv.Atoi(inputIngredient[0])
			reactions[ingredientName].Ingredients[inputIngredient[1]] = quantity
		}
	}
	return reactions
}

func part1(reactions map[string]Reaction, name string, num int) int {

	elements := map[string]int{name: num}
	more := false
	for {
		more = false
		// Look for first element
		for elementName := range elements {
			if elementName != "ORE" && elements[elementName] > 0 {

				more = true
				// Crear
				desired := elements[elementName]
				number := (desired-1)/reactions[elementName].Quantity + 1
				elements[elementName] -= reactions[elementName].Quantity * number

				// Afegir els generadors
				for ingredient := range reactions[elementName].Ingredients {
					elements[ingredient] += reactions[elementName].Ingredients[ingredient] * number
				}
				break
			}
		}

		if more == false {
			break
		}
	}
	return elements["ORE"]
}

// La meva solució, sistema bruteforce ... Tarda una mica en trobar
// el resultat.
func part2Bruteforce(reactions map[string]Reaction, maxOres int) int {

	fuelsCreated := 0
	oresUsed := 0
	for oresUsed < maxOres {
		fuelsCreated++
		oresUsed = part1(reactions, "FUEL", fuelsCreated)
	}
	return fuelsCreated - 1
}

// Aquest l'he trobat entre les solucions de la gent i m'ha semblat brutal ...
func part2(reactions map[string]Reaction, maxOres int) int {
	return (sort.Search(maxOres, func(n int) bool {
		return part1(reactions, "FUEL", n) > maxOres
	}) - 1)
}
