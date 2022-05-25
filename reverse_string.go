package homework

func ReverseString(input string) (output string) {

	var sl []rune = make([]rune, len(input))
	for i, s := range input {
		sl[len(input)-1-i] = s
	}

	output = string(sl)

	return output
}
