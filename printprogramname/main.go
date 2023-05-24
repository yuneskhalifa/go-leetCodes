package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ch := os.Args[0]
	for _, sh := range ch {
		z01.PrintRune(sh)
	}
	z01.PrintRune('\n')
}
