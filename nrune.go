package piscine

func NRune(s string, n int) rune {
s1 := []rune(s)
if n <= 0 || n > len(s1){
	return 0
}
return s1[n-1]
}
