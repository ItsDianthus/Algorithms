package main

import (
	"fmt"
	"strconv"
	"strings"
)

// medium

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	b := strings.Builder{}
	for _, val := range strs {
		b.WriteString(strconv.Itoa(len(val)))
		b.WriteByte('#')
		b.WriteString(val)
	}
	return b.String()
}

func (s *Solution) Decode(st string) []string {
	var res []string
	i := 0
	lenS := 0

	// "4#asdf6#qwerty6#5#%#90"

	for i < len(st) {
		if st[i] == '#' {
			res = append(res, st[i+1:i+1+lenS])
			i = i + 1 + lenS
			lenS = 0
		} else {
			lenS = lenS*10 + int(st[i]-'0')
			i++
		}
	}

	return res
}

func main() {
	//s := "11#1234567891!12#qwertyui..."
	// fmt.Println(s[3:14]) // 2645

	s := Solution{}
	a := s.Encode([]string{"asdf", "qwerty", "5#%#90"})
	fmt.Println(a)
	b := s.Decode(a)
	fmt.Println(b)
}
