package piscine

func Fibonacci(index int) int {
	v := 0
	if index < 0 {
		return -1
	} else if index >= 0 && index < 2 {
		return index
	} else {
		return v + Fibonacci(index-1) + Fibonacci(index-2)
	}
}
