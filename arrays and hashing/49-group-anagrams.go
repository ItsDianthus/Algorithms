package main

import (
	"fmt"
	"slices"
)

// medium

func groupAnagrams(strs []string) [][]string {
	mapa := make(map[string][]string)

	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		mapa[string(b)] = append(mapa[string(b)], str) // Работает и на инит, и если уже было что-то
	}

	var res [][]string // для ускорения можно заменить на вариант ниже (минус аллокации)
	// res := make([][]string, 0, len(mapa))

	for _, val := range mapa {
		res = append(res, val)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"f"}))
}
