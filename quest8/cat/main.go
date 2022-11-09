package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func ReadFile(s string) {
	content, err := os.ReadFile(s)
	if err != nil {
		Print("ERROR: open " + s + ": no such file or directory\n")
		os.Exit(1)
	}
	Print(string(content))
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			ReadFile(arg)
		}
	} else if len(args) == 0 {
		output := os.Stdin
		out, _ := ioutil.ReadAll(output)
		Print(string(out))
	}
}
