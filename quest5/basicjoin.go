package piscine

func BasicJoin(elems []string) string {
	var str string
	for i := range elems {
		str += elems[i]
	}
	return str
}
