package piscine

func TrimAtoi(s string) int {
	sign := 1
	sum := 0
	seen := false
	for _, v := range s {
		if v == '-' && !seen {
			sign = -1
		} else if v >= '0' && v <= '9' {
			sum = sum*10 + int(v-'0')
			seen = true
		}
	}
	return sum * sign
}
