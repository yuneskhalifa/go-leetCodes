package piscine

func CollatzCountdown(start int) int {
	counter := 0
	for i := start; i > 1; i-- {
		if start == 0 || start < 0 {
			return -1
		}
		counter++
	}
	return counter
}
