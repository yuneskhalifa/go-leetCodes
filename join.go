package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i := 0; i < len(strs); i++ {
		s += strs[i]
		if i != len(strs)-1 {
			s += string(sep)
		}
	}
	return s
}
