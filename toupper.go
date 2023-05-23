package piscine

func ToUpper(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r -= 32
		}
		result = append(result, r)
	}
	return string(result)
}
