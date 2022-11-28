package queue

type PriorityQueueDelete struct {
	Elements []int
	Len      int
}

func (q *PriorityQueueDelete) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
	q.Len += 1
}

func (q *PriorityQueueDelete) Dequeue() int {
	if q.Len == 0 {
		return 0
	}
	maxElem := q.Elements[0]
	maxIndex := 0

	for i := 1; i < len(q.Elements); i++ {
		if q.Elements[i] > maxElem {
			maxElem = q.Elements[i]
			maxIndex = i
		}
	}

	if maxIndex != q.Len-1 {
		q.Elements = append(q.Elements[:maxIndex], q.Elements[maxIndex+1:]...)
	} else {
		q.Elements = q.Elements[:q.Len-1]
	}
	return maxElem
}
