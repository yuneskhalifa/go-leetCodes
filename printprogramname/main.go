package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args[0]
	for _, i := range arg[2:] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
