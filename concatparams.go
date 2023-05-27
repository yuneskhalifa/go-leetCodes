package main

func ConcatParams(args []string) string {

	var n string
	for i := 0; i < len(args); i++ {
		n += args[i] + "\n"
	}
	return n
}
