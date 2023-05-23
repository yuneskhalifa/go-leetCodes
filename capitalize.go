package piscine

func Capitalize(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			result = append(result, unicode.ToUpper(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
