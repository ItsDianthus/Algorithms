package main

import "fmt"

// easy
func reverseWords(s string) string {
	var res []byte
	stInd := 0

	for i := 0; i < len(s); i++ {
		if s[i] == byte(' ') {
			for j := i - 1; j >= stInd; j-- {
				res = append(res, s[j])
			}
			stInd = i + 1
			res = append(res, byte(' '))
		}
	}
	if stInd < len(s) {
		for j := len(s) - 1; j >= stInd; j-- {
			res = append(res, s[j])
		}
	}

	return string(res)
}

func main() {
	fmt.Println(reverseWords("Lets getting started "))
}
