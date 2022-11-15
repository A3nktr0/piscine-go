package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{}
	n.Data = data_ref
	n.Next = nil
	if l == nil || l.Data >= n.Data {
		n.Next = l
		return n
	} else {
		current := l
		for current.Next != nil && current.Next.Data < n.Data {
			current = current.Next
		}
		n.Next = current.Next
		current.Next = n
	}
	return l
}
