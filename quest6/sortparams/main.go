package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortArguments(table []string) []string {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}

func main() {
	arguments := os.Args[1:]
	sort_table := SortArguments(arguments)

	for i := range sort_table {
		for _, v := range sort_table[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
