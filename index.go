package piscine

func Index(s, toFind string) int {
	a := len(s)
	suba := len(toFind)
	t := 0
	for i := 0; i < a; i++ {
		j := 0
		t = i
		for j < suba {
			if t < a && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == suba {
			return i
		}
	}
	return -1
}
