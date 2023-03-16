package feature_2

import "errors"

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
}

func (q *Queue) Dequeue() (int, error) {
	if q.size == 0 {
		return -1, errors.New("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item, nil
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return -1, errors.New("empty queue")
	}
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
