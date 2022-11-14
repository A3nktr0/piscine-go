package piscine

func Compact(ptr *[]string) int {
	var clean_str []string
	for _, v := range *ptr {
		if v != "" {
			clean_str = append(clean_str, v)
		}
	}
	*ptr = clean_str
	return len(clean_str)
}
