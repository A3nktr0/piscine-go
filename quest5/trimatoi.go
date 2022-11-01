package piscine

func TrimAtoi(s string) int {
	var num []rune

	for i := range s {
		if IsNumeric(string(s[i])) || s[i] == '-' {
			num = append(num, rune(s[i]))
		}
		// z01.PrintRune(rune(s[i]))
	}

	for i := 0; i < len(num); i++ {
		if num[i] == '-' {
			if i != 0 {
				num = append(num[:i], num[i+1:]...)
			}
		}
	}
	return Atoi(string(num))
}
