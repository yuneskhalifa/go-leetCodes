package piscine

func Capitalize(s string) string {
	var result []byte
	for i, c := range s {
		if unicode.IsLower(c) {
			if i == 0 || !unicode.IsLower(s[i-1]) {
				result = append(result, unicode.ToUpper(c))
			} else {
				result = append(result, c)
			}
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
