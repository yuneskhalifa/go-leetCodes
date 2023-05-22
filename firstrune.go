package piscine

func FirstRune(s string) rune {
	for i, r := range s {
		if r != utf8.RuneError {
			return r
		}
	}

	return 0
}
