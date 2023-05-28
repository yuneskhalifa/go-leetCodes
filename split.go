package piscine

func Split(s, sep string) []string {
	parts := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			parts = append(parts, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	if start < len(s) {
		parts = append(parts, s[start:])
	}
	return parts
}
