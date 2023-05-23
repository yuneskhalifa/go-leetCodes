package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	str := ""
	cap := true

	for index, c := range runes {
		if c >= 'A' && c <= 'Z' {
			runes[index] = c + ('a' - 'A')
		}
	}
	for index, c := range runes {
		if (c >= 'a' && c <= 'z') || ('0' <= c && c <= '9') {
			if cap && (c >= 'a' && c <= 'z') {
				runes[index] = c - ('a' - 'A')
			}
			cap = false
		} else {
			cap = true
		}
		str += string(runes[index])
	}
	return str
}
