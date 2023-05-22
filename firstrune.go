package piscine

import "unicode/utf8"

func FirstRune(s string) rune {
	r, size := utf8.DecodeRuneInString(s)
	return r
}
