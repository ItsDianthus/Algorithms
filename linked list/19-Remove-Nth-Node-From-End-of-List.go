package main

import "fmt"

// medium
// а это точно был медиум?

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	cnt := 0

	for p1 != nil {
		cnt += 1
		p1 = p1.Next
	}

	p1 = head
	last := head
	for i := 0; i < cnt-n; i++ {
		last = p1
		p1 = p1.Next
	}
	// fmt.Println(p1.Val) // n с конца

	if cnt-n == 0 {
		//fmt.Println("first!")
		return p1.Next
	} else if n == 1 {
		last.Next = nil
		return head
	} else {
		last.Next = p1.Next
		return head
	}
}

func printNodes(head *ListNode) {
	p1 := head

	for p1 != nil {
		fmt.Print(p1.Val, " ")
		p1 = p1.Next
	}

}

func main() {
	// Создаём список a1 из 1 узла
	_ = &ListNode{Val: 10, Next: nil}

	// Создаём список a2 из 3 узлов
	a2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	printNodes(removeNthFromEnd(a2, 2))

}
