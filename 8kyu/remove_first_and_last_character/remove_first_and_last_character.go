package remove_first_and_last_character

func RemoveChar(word string) string {
	return word[1 : len(word)-1]

	// my first solution
	//runes := []rune(word)
	//return string(runes[1 : len(runes)-1])
}
