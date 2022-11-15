package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil && l2.Head == nil {
		return
	}
	if l1.Head == nil {
		l1.Head = l2.Head
	} else {
		current := l1.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = l2.Head
	}
}
