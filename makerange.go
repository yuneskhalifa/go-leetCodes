package piscine

func MakeRange(min int, max int) []int {
	var s []int
	if min >= max {
		return s
	}
	var x = make([]int, max-min)
	counter := 0
	for i := min; i < max; i++ {
		x[counter] = i
		counter++
	}
	return x
}
