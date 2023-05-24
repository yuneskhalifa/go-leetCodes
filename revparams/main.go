package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	arglen := len(arg) - 1
	for i := arglen; i >= 0; i-- {
		for _, j := range arg[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
