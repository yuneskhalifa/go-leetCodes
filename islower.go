package piscine

func IsLower(s string) bool {
	func isLowercase(s string) bool {
		for _, r := range s {
		  if r < 'a' || r > 'z' {
			return false
		  }
		}
		return true
	  }
}
		
