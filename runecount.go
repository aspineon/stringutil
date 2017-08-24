package stringutil

func RuneCount(s string) map[rune]int {
	r := []rune(s)
	m := make(map[rune]int)
	for i := 0; i < len(r); i = i+1 {
		m[r[i]] = m[r[i]] + 1
	}

	return m

}