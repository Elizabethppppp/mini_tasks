package slices

type Queue struct {
	Data []int
}

func (q *Queue) Enqueue(v int) {
	q.Data = append(q.Data, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.Data) == 0 {
		return 0, false
	}
	v := q.Data[0]
	q.Data = q.Data[1:]
	return v, true
}
