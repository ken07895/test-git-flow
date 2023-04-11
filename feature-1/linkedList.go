package feature_1

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func NewHead(Val int) *LinkedList {
	head := LinkedList{Val: Val}
	fmt.Println("test mond 3")
	return &head
}

func Search(l *LinkedList, searchingVal int) *LinkedList {
	l = l.Next
	for l.Next != nil {
		if l.Val == searchingVal {
			return l
		}
		l = l.Next
	}
	return nil
}

func AppendList(head *LinkedList, val int) *LinkedList {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &LinkedList{Val: val}
	return head
}

// add comment
func FindingMidList(head *LinkedList) *LinkedList {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Println(slow.Val)
	fmt.Println(fast.Val)
	return slow
}

func ShowAllValues(head *LinkedList) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("eiei za")
}
