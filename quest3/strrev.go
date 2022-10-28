package piscine

func StrRev(s string) string {
	sr := []byte(s)
	for i, j := 0, len(sr)-1; i < j; i, j = i+1, j-1 {
		sr[i], sr[j] = sr[j], sr[i]
	}
	return string(sr)
}
