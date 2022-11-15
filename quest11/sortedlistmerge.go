package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	merged := &NodeI{}
	n := merged
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}
		n = n.Next
	}
	if n1 != nil {
		n.Next = n1
	} else {
		n.Next = n2
	}
	return merged.Next
}
