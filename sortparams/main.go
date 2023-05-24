package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i, a := range arg {
		for j, b := range arg {
			if b > a {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for _, c := range arg {
		for _, d := range c {
			z01.PrintRune(d)
		}
		z01.PrintRune('\n')
	}
}
