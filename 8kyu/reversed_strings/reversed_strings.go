package reversed_strings

func Reverse(s string) string {
	runes := []rune(s)
	for i := len(runes); i < 0; i-- {
	}
	return s
}
