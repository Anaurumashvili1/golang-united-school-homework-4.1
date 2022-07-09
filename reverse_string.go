package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	byteString := []rune(input)
	for i, j := 0, len(byteString)-1; i < j; i, j = i+1, j-1 {
		byteString[i], byteString[j] = byteString[j], byteString[i]
	}
	return string(byteString)
}
