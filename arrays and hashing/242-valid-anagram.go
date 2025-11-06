package main

import (
	"fmt"
)

// easy

func isAnagram(s string, t string) bool {
	r1 := []rune(s)
	r2 := []rune(t)

	myMap := make(map[rune]int)

	if len(r1) != len(r2) {
		return false
	}

	for i := 0; i < len(r1); i++ {
		if _, exist := myMap[r1[i]]; exist {
			myMap[r1[i]] += 1
		} else {
			myMap[r1[i]] = 1
		}
	}
	for i := 0; i < len(r1); i++ {
		if _, exist := myMap[r2[i]]; exist {
			myMap[r2[i]] -= 1 // если существует
			if myMap[r2[i]] == 0 {
				delete(myMap, r2[i])
			}
		} else {
			return false
		}
	}

	return true
}

// Попроще: держим не мапу, а массив на 26 символов
// Побеждает по памяти и времени
func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	for _, c := range t {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("", ""))
}
