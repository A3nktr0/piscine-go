package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func Print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNbr(n int) {
	nb := '0'
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n > 9 {
		PrintNbr(n / 10)
		n %= 10
	}
	if n > 0 && n <= 9 {
		for i := 1; i <= n; i++ {
			nb++
		}
	}
	z01.PrintRune(nb)
}

func SetPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	SetPoint(points)
	Print("x = ")
	PrintNbr(points.x)
	Print(", y = ")
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
