package piscine

import "github.com/01-edu/z01"

func ccc() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; i < '9'; i++ {
			for k := '0'; k < '9'; k++ {
				for l := '1'; l < '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
