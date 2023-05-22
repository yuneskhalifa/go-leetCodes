package piscine

func LastRune(s string) rune {
	nh := []rune(s)
	n := len(nh) - 1
	return rune(s[n])
}
