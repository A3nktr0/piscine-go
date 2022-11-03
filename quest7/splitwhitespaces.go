package piscine

func SplitWhiteSpaces(s string) []string {
	var str string
	var str_tab []string
	for i := range s {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			str += string(s[i])
		} else {
			if len(str) != 0 {
				str_tab = append(str_tab, str)
				str = ""
			}
		}
	}
	if len(str) != 0 {
		str_tab = append(str_tab, str)
	}
	return str_tab
}
