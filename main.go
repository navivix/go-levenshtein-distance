package main

import "fmt"

func min(is ...uint) uint {
	min := is[0]
	for _, i := range is[1:] {
		if i < min {
			min = i
		}
	}
	return min
}

func levenshtein(first, second string) [][]uint {
	x := len(first) + 1
	y := len(second) + 1

	ret := make([][]uint, y)

	for i := range ret {
		ret[i] = make([]uint, x)
	}

	for i := range ret[0] {
		ret[0][i] = uint(i)
	}

	for i := range ret {
		ret[i][0] = uint(i)
	}

	for i := 1; i < y; i++ {
		for j := 1; j < x; j++ {
			if first[j-1] == second[i-1] {
				ret[i][j] = ret[i-1][j-1]
			} else {
				ret[i][j] = min(ret[i][j-1], ret[i-1][j-1], ret[i-1][j]) + 1
			}
		}
	}

	return ret
}

func display(first, second string, lev [][]uint) {
	fmt.Println("\nLevenshtein matrix:")

	fmt.Print("    ")

	for i := range first {
		fmt.Print(string(first[i]), " ")
	}

	fmt.Print("\n  ")

	for i := range lev[0] {
		fmt.Print(lev[0][i], " ")
	}

	fmt.Println("")

	for i := 1; i < len(lev); i++ {
		fmt.Print(string(second[i-1]), " ")
		for j := range lev[i] {
			fmt.Print(lev[i][j], " ")
		}
		fmt.Println("")
	}

	fmt.Println("\nDistance:", lev[len(lev)-1][len(lev[0])-1])
}

func main() {
	var first, second string

	fmt.Print("First word: ")
	fmt.Scanln(&first)

	fmt.Print("Second word: ")
	fmt.Scanln(&second)

	ret := levenshtein(first, second)

	display(first, second, ret)
}
