package main

// easy

func intersect(nums1 []int, nums2 []int) []int {
	nums1, nums2 = arrAscending(nums1, nums2)
	myMap := make(map[int]int)
	res := make([]int, 0, len(nums1))

	for i := 0; i < len(nums1); i++ {
		myMap[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if len(myMap) == 0 {
			return res
		}
		if _, exist := myMap[nums2[i]]; exist {
			myMap[nums2[i]]--
			if myMap[nums2[i]] == 0 {
				delete(myMap, nums2[i])
			}
			res = append(res, nums2[i])
		}
	}
	return res
}

func arrAscending(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums2) < len(nums1) {
		return nums2, nums1
	} else {
		return nums1, nums2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 5}

	nums1, nums2 = arrAscending(nums1, nums2)

}
