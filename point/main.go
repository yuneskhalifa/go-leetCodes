package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(nbr int) {
	x := '0'
	if nbr/10 > 0 {
		PrintInt(nbr / 10)
	}
	for i := 0; i < nbr%10; i++ {
		x++
	}
	z01.PrintRune(x)
}
func converToString(str string) {
	s := []rune(str)
	for index := range s {
		z01.PrintRune(s[index])
	}
}

func main() {
	var points point
	setPoint(&points)
	converToString("x = ")
	PrintInt(points.x)
	converToString(", y = ")
	PrintInt(points.y)
	z01.PrintRune('\n')
}
