package main

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func Match(s, subs string, i int) bool {
	for j := 0; j < len(subs); j++ {
		if s[i] != subs[j] {
			return false
		}
		i++
	}
	return true
}

func CheckContain(s, subs string) bool {
	for i := 0; i < len(s)-len(subs); i++ {
		if Match(s, subs, i) {
			return true
		}
	}
	return false
}

func Insert(s string) string {
	var insert []string
	output := ""
	for i := range s {
		if s[i] == '=' {
			insert = append(insert, s[i+1:])
		}
	}
	for _, v := range insert {
		output += v
	}
	return output
}

func Order(s string) string {
	table := []rune(s)
	var output string
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	for _, v := range table {
		output += string(v)
	}
	return output
}

func main() {
	arguments := os.Args[1:]
	var str string
	var str_insert string
	is_order := false
	is_help := false

	for _, v := range arguments {
		if CheckContain(v, "-i") || CheckContain(v, "--index") {
			str_insert += Insert(v)
		} else if v == "-o" || v == "--order" {
			is_order = true
		} else if arguments[0] == "-h" || arguments[0] == "--help" {
			is_help = true
		} else {
			str += v
		}
	}

	if is_order {
		str = Order(str + str_insert)
	} else {
		str += str_insert
	}
	if len(str) >= 1 {
		fmt.Println(str)
	} else if len(arguments) == 0 || is_help {
		Help()
	}
}
