package kata

import "unicode"

func GetCount(str string) (count int) {
	for _, char := range str {
		switch unicode.ToLower(char) {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
