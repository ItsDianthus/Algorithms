package main

// easy
// они соединятся поэтому верим

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p1_fast, p2 := head, head

	for p1_fast != nil || p1_fast.Next != nil {
		p1_fast = p1_fast.Next.Next
		p2 = p2.Next
		if p1_fast == p2 {
			return true
		}
	}
	return false
}

// Пример использования
func main() {

}
