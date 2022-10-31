package piscine

func Fibonacci(index int) int {
	v := 0
	if index == 0 {
		return v
	} else if index == 1 {
		return v + 1
	} else if index >= 2 {
		return v + Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return -1
	}
}
