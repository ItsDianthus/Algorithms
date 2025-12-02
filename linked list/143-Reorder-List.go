package main

// medium
// идея конечно хард
// 1) бежим до середины (меньшую середину забираем)
// 2) вторую половину переворачиваем
// 3) объединяем два линкд листа
// в третьем пункте трудновато но в целом норм

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// 1)
	if head == nil || head.Next == nil {
		return
	}
	p1 := head
	p2 := head.Next

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	// теперь p1.Next это хэд нашего листа который мы переворачиваем
	// 2)
	curr := p1.Next
	p1.Next = nil
	var last *ListNode = nil
	next := curr.Next
	for {
		curr.Next = last
		last = curr
		curr = next
		if curr == nil {
			break
		}
		next = curr.Next
	}
	// last это head нового отрезка

	p1 = head
	p2 = last
	for p1 != nil {
		next1 := p1.Next
		next2 := p2.Next

		p1.Next = p2
		// проверяем, есть ли куда ссылаться после p2
		if next1 == nil {
			break
		}
		p2.Next = next1

		p1 = next1
		p2 = next2
	}
}

func main() {

}
