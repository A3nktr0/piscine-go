package piscine

import (
	"github.com/01-edu/z01"
)

func ConvertToBinary(nb int) {
	var bin []int

	for nb != 0 {
		bin = append(bin, nb%2)
		nb = nb / 2
	}
	for i := len(bin) - 1; i >= 0; i-- {
		PrintNbr(bin[i])
	}
}

func ConvertToHexa(nb int) {
	vh := []rune{
		'1', '2', '3', '4', '5', '6',
		'7', '8', '9', 'A', 'B', 'C',
		'D', 'E', 'F',
	}
	var hex []int

	for nb != 0 {
		hex = append(hex, nb%16)
		nb = nb / 16
	}
	for i := len(hex) - 1; i >= 0; i-- {
		for j := 0; j < len(vh); j++ {
			if j == hex[i] {
				z01.PrintRune(vh[j] - 1)
			}
		}
	}
}

func ConvertToChoumi(nb int) {
	vc := []rune{'c', 'h', 'o', 'u', 'm', 'i'}
	var choumi []int

	for nb != 0 {
		choumi = append(choumi, nb%6)
		nb = nb / 6
	}
	for i := len(choumi) - 1; i >= 0; i-- {
		for j := 0; j < len(vc); j++ {
			if j == choumi[i] {
				z01.PrintRune(vc[j])
			}
		}
	}
}

func PrintNbrBase(nbr int, base string) {
	switch {
	case base == "0123456789":
		PrintNbr(nbr)
	case base == "01":
		if nbr < 0 {
			z01.PrintRune('-')
			ConvertToBinary(nbr * -1)
		} else {
			ConvertToBinary(nbr)
		}
	case base == "0123456789ABCDEF":
		if nbr < 0 {
			z01.PrintRune('-')
			ConvertToHexa(nbr * -1)
		} else {
			ConvertToHexa(nbr)
		}
	case base == "choumi":
		{
			if nbr < 0 {
				z01.PrintRune('-')
				ConvertToChoumi(nbr * -1)
			} else {
				ConvertToChoumi(nbr)
			}
		}
	default:
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
