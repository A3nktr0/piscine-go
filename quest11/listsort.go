package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	current := l
	for current != nil {
		next := current.Next
		for next != nil {
			if current.Data > next.Data {
				current.Data, next.Data = next.Data, current.Data
			}
			next = next.Next
		}
		current = current.Next
	}
	return l
}
