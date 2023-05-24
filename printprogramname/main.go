package main

import (
	"os"

	"github.com/01-edu/z01"

	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
