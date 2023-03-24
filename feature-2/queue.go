package feature_2

import (
	"errors"
	"fmt"
)

// add comment
type Queue struct {
	items []int
	size  int
}

func Construct(maxSize int) Queue {
	qItems := make([]int, 0, maxSize)
	return Queue{qItems, 0}
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	q.size++
	fmt.Println("mond 1")
	fmt.Println("mond 2")
	fmt.Println("Kai you")
	fmt.Println("branch 1")
	fmt.Println("branch 1 1")
}

func (q *Queue) Dequeue() (int, error) {
	if q.size == 0 {
		return -1, errors.New("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	fmt.Println("mond 2")
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return -1, errors.New("empty queue")
	}
	fmt.Println("Kai you")
	return q.items[0], nil
}

func (q *Queue) Search(val int) int {
	for idx_searchingConflict, searchingVal := range q.items {
		if val == searchingVal {
			return idx_searchingConflict
		}
	}
	return -1
}

func (q *Queue) PrintItems() {
	fmt.Println(q.items)
}

func (q *Queue) PrintItem(idx int) {
	if idx >= q.size {
		fmt.Println("failed")
	}
	fmt.Println(q.items[idx])
}

func (q Queue) Test() {
	fmt.Println("comment 1")
	fmt.Println("comment 2")
	fmt.Println("comment 5")
	fmt.Println("mond 4")
	fmt.Println("mond 5")
}
