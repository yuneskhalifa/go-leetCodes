package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("Too many arguments\n")
	} else if len(os.Args) == 1 {
		fmt.Printf("File name missing\n")
	} else {
		fmt.Printf("Almost there!!\n")
	}
}
