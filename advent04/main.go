package main

import "fmt"

func main() {

	start := 137683
	stop := 596253

	valids1 := 0
	valids2 := 0

	for candidate := start; candidate <= stop; candidate++ {
		if isValid(candidate) {
			valids1++
		}
		if isValid2(candidate) {
			valids2++
		}
	}

	fmt.Printf("589999 %t\n", isValid2(589999))
	fmt.Printf("112233 %t\n", isValid2(112233))
	fmt.Printf("123444 %t\n", isValid2(123444))
	fmt.Printf("111122 %t\n", isValid2(111122))

	fmt.Printf("PART 1 Els possibles passwords són %d\n", valids1)
	fmt.Printf("PART 2 Els possibles passwords són %d\n", valids2)
}

func isValid(candidate int) bool {

	digitAnterior := 9999
	digitRepetit := false

	for i := 1; i <= 100000; i = i * 10 {
		digit := (candidate / i) % 10
		if digit == digitAnterior {

			digitRepetit = true
		}
		if digit > digitAnterior {
			return false
		}
		digitAnterior = digit
	}
	if digitRepetit {
		fmt.Printf("Candidat: %d\n", candidate)
	}
	return digitRepetit
}

func isValid2(candidate int) bool {

	digitAnterior := 9999
	digitRepetit := 0
	repetitDosCops := false

	for i := 1; i <= 100000; i = i * 10 {
		digit := (candidate / i) % 10
		if digit == digitAnterior {

			digitRepetit = digitRepetit + 1

		} else {
			if digitRepetit == 1 {
				repetitDosCops = true
			}
			digitRepetit = 0
		}
		if digit > digitAnterior {
			return false
		}
		digitAnterior = digit
	}
	if repetitDosCops || digitRepetit == 1 {
		fmt.Printf("Candidat: %d\n", candidate)
	}
	return repetitDosCops || digitRepetit == 1
}
