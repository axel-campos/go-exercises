package main

import (
	"lists.com/sixteen"
)

func allCombinations() {
	goal := 100
	universe := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	char_combinations := []rune{'_', '-', '+'}

	//posible_combinations := cartesianProductN(char_combinations, 8)
	posible_combinations := cartesianProductN(char_combinations, 8)

	/* for _, set := range posible_combinations {
		print("[")
		for _, value := range set {
			print(" ", string(value), " ")
		}
		println("]")
	} */
	println(posible_combinations)

	for _, combination := range posible_combinations {
		if testCombination(universe, combination, goal) {
			printCombination(universe, combination, goal)
		}
	}
	/* basic_combination := []rune{'+', '+', '+', '+', '+', '+', '+', '+'}
	printCombination(universe, basic_combination, goal) */
}

func cartesianProductN(set []rune, times int) (result_set [][]rune) {
	if times == 0 {
		return [][]rune{nil}
	}
	acum_set := cartesianProductN(set, times-1)
	for _, rrune := range set {
		for _, sets := range acum_set {
			result_set = append(result_set, append([]rune{rrune}, sets...))
		}
	}
	return
}

func testCombination(universe []int, combinations []rune, goal int) bool {
	total := 0
	var last_operation int = -1
	combinations = append([]rune{'+'}, combinations...)
	if len(universe) == len(combinations) {
		for i := 0; i < len(universe); i++ {
			if combinations[i] == '_' {
				continue
			} else {
				if last_operation >= 0 {
					if combinations[last_operation] == '+' {
						total += sixteen.BaseToDecimal(universe[last_operation:i], 10)
					} else {
						total -= sixteen.BaseToDecimal(universe[last_operation:i], 10)
					}
					//println("Total: ", total)
					//println("Slice value: ", sixteen.BaseToDecimal(universe[last_operation:i], 10))
				}
				last_operation = i
			}
		}
		if combinations[last_operation] == '+' {
			total += sixteen.BaseToDecimal(universe[last_operation:], 10)
		} else {
			total -= sixteen.BaseToDecimal(universe[last_operation:], 10)
		}
	}
	//println("Total: ", total)
	if total == goal {
		return true
	} else {
		return false
	}
}

func printCombination(universe []int, combinations []rune, goal int) {
	for i := 0; i < len(universe); i++ {
		print(universe[i])
		if i < len(combinations) {
			if combinations[i] != '_' {
				print(string(combinations[i]))
			}
		}
	}
	println(" =", goal)
}

func main() {
	allCombinations()
}
