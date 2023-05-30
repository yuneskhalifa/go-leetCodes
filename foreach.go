package piscine

import "fmt"

func PrintNbr(b int) {
	fmt.Print(b)
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}
