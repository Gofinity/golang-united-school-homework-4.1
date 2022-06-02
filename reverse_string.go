package reverse_string

func ReverseString(input string) (output string) {

	inputRunes := []rune(input)
	for i := len(inputRunes) - 1; i >= 0; i-- {
		output += string(inputRunes[i])
	}
	return output
}
