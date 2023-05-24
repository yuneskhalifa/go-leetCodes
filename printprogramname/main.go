package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := []rune(os.Args[0])
	for i := 0; i < len(name); i++ {
		z01.PrintRune(name[i])
	}
}
