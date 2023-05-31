package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		return
	}
	for _, jh := range arguments[0] {
		if !(jh >= '0' && jh <= '9') {
			return
		}
	}
	for _, jh := range arguments[2] {
		if !(jh >= '0' && jh <= '9') {
			return
		}
	}
	if Atoi(arguments[0]) <= -9223372036854775807 || Atoi(arguments[0]) >= 9223372036854775807 {
		return
	}
	if Atoi(arguments[2]) <= -9223372036854775807 || Atoi(arguments[2]) >= 9223372036854775807 {
		return
	}
	var numb1 int = Atoi(arguments[0])
	var numb2 int = Atoi(arguments[2])
	opr := arguments[1]
	if opr == "+" {
		fmt.Println(fmt.Sprint(numb1 + numb2))
		return
	} else if opr == "-" {
		fmt.Println(fmt.Sprint(numb1 - numb2))
		return
	} else if opr == "*" {
		fmt.Println(fmt.Sprint(numb1 * numb2))
		return
	} else if opr == "/" {
		if numb2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(fmt.Sprint(numb1 / numb2))
		return
	} else if opr == "%" {
		if numb2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(fmt.Sprint(numb1 % numb2))
		return
	} else {
		return
	}
}

func Atoi(s string) int {
	runes := []rune(s)
	nb := 0
	isNeg := 1
	for i, c := range runes {
		if c == '-' {
			if i == 0 {
				isNeg = -1
			} else {
				isNeg = 1
				nb = 0
				break
			}
		} else if c == '+' {
			isNeg = 1
			if i != 0 {
				nb = 0
				break
			}
		} else if c >= '0' && c <= '9' {
			nb = (nb*10 + int(c) - '0')
		} else {
			isNeg = 1
			nb = 0
			break
		}
	}
	return nb * isNeg
}
