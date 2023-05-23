package piscine

func IsLower(s string) bool {
	for _, sh := range s {
		if sh >= 'a' || sh <= 'z' {
			return true
		} else {
			return false
		}
	}

}
