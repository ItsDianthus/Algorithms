package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var fin []rune

	for _, r := range s {
		// Можно было бы использовать для большего функционала обработки,
		// А так то в нашем случае достаточно одного ифа
		switch {
		case unicode.IsLetter(r):
			fin = append(fin, r)
		case unicode.IsDigit(r):
			fin = append(fin, r)
		default:
			// skip
		}
	}

	// офигеть можно было два индекса в один фор
	for i, j := 0, len(fin)-1; i < j; i, j = i+1, j-1 {
		if fin[i] != fin[j] {
			return false
		}
	}
	return true
}

func main() {
	isPalindrome("A s,,,, dsjss")

}
