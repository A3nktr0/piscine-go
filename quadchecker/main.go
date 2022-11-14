package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsQuad(str string, x, y int, a, b, c, d, tb, s rune) bool {
	var quad_output []rune
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				if j == 0 {
					quad_output = append(quad_output, a)
				} else if j == x-1 {
					quad_output = append(quad_output, b)
				} else {
					quad_output = append(quad_output, tb)
				}
			} else if i == y-1 {
				if j == 0 {
					quad_output = append(quad_output, c)
				} else if j == x-1 {
					quad_output = append(quad_output, d)
				} else {
					quad_output = append(quad_output, tb)
				}
			} else {
				if j == 0 || j == x-1 {
					quad_output = append(quad_output, s)
				} else {
					quad_output = append(quad_output, ' ')
				}
			}
		}
		quad_output = append(quad_output, '\n')
	}
	quad_str := string(quad_output)
	return str == quad_str
}

func main() {
	var output []rune
	x, y, nb := 0, 0, 0
	flag := true

	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, r)
	}
	str := string(output)

	for _, r := range output {
		if r != '\n' && flag {
			x++
		}
		if r == '\n' {
			y++
			flag = false
		}
	}

	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	} else {
		switch {
		case IsQuad(str, x, y, 'o', 'o', 'o', 'o', '-', '|'):
			fmt.Printf("[quadA] [%d] [%d]\n", x, y)
			return
		case IsQuad(str, x, y, '/', '\\', '\\', '/', '*', '*'):
			fmt.Printf("[quadB] [%d] [%d]\n", x, y)
			return
		}
		if IsQuad(str, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
			nb++
			fmt.Printf("[quadC] [%d] [%d]", x, y)
		}
		if IsQuad(str, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
			if nb > 0 {
				fmt.Print(" || ")
			}
			nb++
			fmt.Printf("[quadD] [%d] [%d]", x, y)
		}
		if IsQuad(str, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
			if nb > 0 {
				fmt.Print(" || ")
			}
			nb++
			fmt.Printf("[quadE] [%d] [%d]", x, y)
		}
	}
	if nb > 0 {
		fmt.Println()
		return
	}
	fmt.Println("Not a Raid function")
}
