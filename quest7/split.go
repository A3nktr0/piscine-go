package piscine

func Split(s, sep string) []string {
	var str_tab []string
	s_sep := 0
	e_sep := s_sep + len(sep)

	for i := 0; i < len(s)-len(sep); {
		for j := i; j < len(s)-len(sep); j++ {
			if s[j:j+len(sep)] == sep {
				str_tab = append(str_tab, s[i:j])
				s_sep = j
				i = j + len(sep)
			}
		}
		i++
	}
	str_tab = append(str_tab, s[s_sep+e_sep:])
	return str_tab
}
