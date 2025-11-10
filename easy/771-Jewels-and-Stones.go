package main

// easy

func numJewelsInStones(jewels string, stones string) int {
	cntRes := 0
	set := make(map[byte]bool)

	for i := 0; i < len(jewels); i++ {
		set[jewels[i]] = true
	}

	for i := 0; i < len(stones); i++ {
		if set[stones[i]] {
			cntRes++
		}
	}

	return cntRes
}

func main() {
}
