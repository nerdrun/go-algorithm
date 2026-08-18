package disemvowel_trolls

import "strings"

func Disemvowel(comment string) string {
	result := ""
	for i := 0; i < len(comment); i++ {
		char := string(comment[i])
		lower := strings.ToLower(char)
		if lower != "a" && lower != "e" && lower != "i" && lower != "o" && lower != "u" {
			result += char
		}
	}
	return result

	// Love this one
	//result := ""
	//for _, v := range comment {
	//	switch v {
	//	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
	//		continue
	//	}
	//	result += string(v)
	//}
	//return result
}
