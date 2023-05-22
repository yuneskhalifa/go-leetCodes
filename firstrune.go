package piscine

func FirstRune(s string) rune {
	r, size := utf8.DecodeRuneInString(s)
	return r
}
