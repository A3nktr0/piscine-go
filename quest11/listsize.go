package piscine

func ListSize(l *List) int {
	count := 1
	if l.Head == nil {
		return 0
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
			count++
		}
	}
	return count
}
