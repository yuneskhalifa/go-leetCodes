package piscine

func Capitalize(s string) string {
	var result []byte
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			if i == 0 || s[i-1] < 'a' || s[i-1] > 'z' {
				result = append(result, byte(c-32))
			} else {
				result = append(result, c)
			}
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
