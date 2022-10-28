package piscine

func UltimateDivMod(a *int, b *int) {
	tempa := *a
	tempb := *b

	*a = tempa / tempb
	*b = tempa % tempb
}
