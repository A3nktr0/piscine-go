package piscine

func MakeRange(min, max int) []int {
	size := 0
	var array []int

	if min < max {
		for i := min; i < max; i++ {
			size++
		}

		array = make([]int, size)

		for i := 0; i < size; i++ {
			array[i] = min + i
		}
	}
	return array
}
