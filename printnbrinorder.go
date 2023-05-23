package piscine

import "fmt"

func PrintNbrInOrder(n int) {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				var temp int
				temp = digits[i]
				digits[i] = digits[j]
				digits[j] = temp
			}
		}
	}
	for _, d := range digits {
		fmt.Print(d)
	}
}
