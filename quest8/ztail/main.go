package main

import (
	"fmt"
	"os"
)

// func Print(s string) {
// 	for _, r := range s {
// 		z01.PrintRune(r)
// 	}
// }

func ConvertToNumber(n rune) int {
	nb := 0
	if n >= '0' && n <= '9' {
		nb = int(n) - 48
	}
	return nb
}

func Atoi(s string) int {
	nb := 0

	for i, v := range s {
		if i != len(s)-1 {
			nb *= 10
			nb += ConvertToNumber(v) * 10
		} else {
			nb += ConvertToNumber(v)
		}
	}
	return nb
}

func main() {
	args := os.Args[1:]
	n := Atoi(args[1])
	c := 0
	exit_status := 0
	if args[0] == "-c" {
		for _, arg := range args[2:] {
			output, err := os.ReadFile(arg)
			if err != nil {
				c++
				fmt.Printf("open %s: no such file or directory\n", arg)
				exit_status = 1
			} else {

				lenArg := len(output) - n
				if lenArg < 0 {
					lenArg = 0
				}
				if c == 0 {
					fmt.Printf("==> %s <==\n%v", arg, string(output[lenArg:]))
					c++
				} else if c >= 1 {
					fmt.Printf("\n==> %s <==\n%v", arg, string(output[lenArg:]))
				}
			}
		}
		if exit_status != 0 {
			os.Exit(1)
		}
	}
}
