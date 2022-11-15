package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	for i := 0; i < pos; i++ {
		current = current.Next
		if current == nil {
			return nil
		}
	}
	return current
}
