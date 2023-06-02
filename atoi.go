package piscine

func Atoi(s string) int {
	runes := []rune(s)
	var result int
	var neg int = 1
	for i, c := range runes {
		if i == 0 && c == '-' {
			neg = -1
		} else if i == 0 && c == '+' {
			continue
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			result = 0
			break
		}
	}
	return result * neg
}
