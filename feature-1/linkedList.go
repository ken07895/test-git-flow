package feature_1

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func NewHead(Val int) *LinkedList {
	head := LinkedList{Val: Val}
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
	return slow
}
