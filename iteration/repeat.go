package iteration

import (
	"strings"
)

// Repeat returns character repeated 5 times
func Repeat(character string, quantity int) string  {
	var repeated string

	for i := 0; i < quantity; i++ {
		repeated += character
	}

	return repeated
}

// LetterEqual returns amount of letter occurrence in a text
func LetterEqual(text, letter string) (count int) {
	countText := len([]rune(text))
	text = LowerCase(text)

	for i := 0; i < countText; i++ {
		result := strings.Compare(string(text[i]), letter)

		if result == 0 {
			count += 1
		}
	}

	return
}

// LowerCase transform characters in lower case
func LowerCase(s string) string {
	return strings.ToLower(s)
}