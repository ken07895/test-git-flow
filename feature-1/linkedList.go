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

func AppendList(head *LinkedList, val int) {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &LinkedList{Val: val}
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

func AppendLists(head *LinkedList, vals []int) {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	for _, val := range vals {
		cur.Next = &LinkedList{Val: val}
		cur = cur.Next
	}
}
