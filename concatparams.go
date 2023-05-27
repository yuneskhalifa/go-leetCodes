package piscine

func ConcatParams(args []string) string {
	var a string
	for index, c := range args {
		a += c
		if index != len(args)-1 {
			a += string('\n')
		}
	}
	return a
}
