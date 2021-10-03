package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	r := q.items[0]
	q.items = q.items[1:]
	return r
}
