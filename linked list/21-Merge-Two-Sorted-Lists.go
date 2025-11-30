package main

// easy
// но гпт поругался что в начале надо инитить по другому хз

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		}
		return list1
	}
	first, l1, l2 := MaxList(list1, list2)
	res := first
	var curr *ListNode = nil

	for l1 != nil && l2 != nil {
		curr, l1, l2 = MaxList(l1, l2)
		first.Next = curr
		first = curr
	}
	if l1 == nil {
		first.Next = l2
	} else if l2 == nil {
		first.Next = l1
	}

	return res
}

// 1 -> 3 -> 5
// 1 -> 2

// first: 1
// list1: 1->3->5
// list2: 2

// first: 1->1
// list1: 3->5
// list2: 2
// curr: 1

func MaxList(list1 *ListNode, list2 *ListNode) (min, l1, l2 *ListNode) {
	if list1.Val < list2.Val {
		return list1, list1.Next, list2
	}
	return list2, list2.Next, list1
}

// nil (1, 2, 3, 4, 5)

// Пример использования
func main() {

}
