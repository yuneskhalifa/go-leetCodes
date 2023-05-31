package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}
	result := ""
	for i, ch := range str {
		if (i+1)%3 != 0 {
			result += string(ch)
		}
	}
	return result + "\n"
}
