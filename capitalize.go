package piscine

func Capitalize(s string) string {
	var result []rune
	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			if i == 0 || s[i-1] < 'a' || s[i-1] > 'z' {
				result = append(result, unicode.ToUpper(r))
			} else {
				result = append(result, r)
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
