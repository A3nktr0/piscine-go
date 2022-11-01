package piscine

func AlphaCount(s string) int {
	alpha := "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	count := 0

	for _, v := range s {
		for _, w := range alpha {
			if v == w {
				count++
			}
		}
	}
	return count
}

// func AlphaCount(s string) int {
//     minMaj := 65
//     maxMaj := 90
//     minMin := 97
//     maxMin := 122
//     count := 0

//     for _, v := range s {
//         if (v >= rune(minMaj) && v <= rune(maxMaj)) || (v >= rune(minMin) && v <= rune(maxMin)) {
//             count++
//         }
//     }
//     return count
// }
