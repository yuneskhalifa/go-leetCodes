package piscine

import "strings"

func SplitWhiteSpaces(s string) []string {
	tab := strings.Fields(s)
	return tab
}
