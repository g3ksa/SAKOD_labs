package queue

type Queue struct {
	Elements []int
	Len      int
}

func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
	q.Len += 1
}

func (q *Queue) Dequeue() int {
	if q.Len == 0 {
		return 0
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	q.Len -= 1
	return firstElement
}
