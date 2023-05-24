package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[j] < arg[i] {
				temp := arg[i]
				arg[i] = arg[j]
				arg[j] = temp
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
