package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	s_name := string(name[0])
	var t_name []string

	for i := range s_name {
		if s_name[i] == '/' {
			t_name = append(t_name, s_name[i+1:])
		}
	}
	ret := string(t_name[len(t_name)-1])

	for _, v := range ret {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
