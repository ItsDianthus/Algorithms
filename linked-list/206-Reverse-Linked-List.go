package main

// easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode = nil
	curr := head
	if head == nil {
		return nil
	}
	for curr != nil {
		next := curr.Next // next 2
		curr.Next = last  // 1->nil
		last = curr       // last = 1
		curr = next       // curr = 2
	}
	return last
}

// nil (1, 2, 3, 4, 5)

// Пример использования
func main() {

}
