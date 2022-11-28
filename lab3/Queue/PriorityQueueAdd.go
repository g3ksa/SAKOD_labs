package queue

type PriorityQueueAdd struct {
	Elements []int
	Len      int
}

func (q *PriorityQueueAdd) Enqueue(elem int) {
	if q.Len == 0 {
		q.Elements = append(q.Elements, elem)
	} else {
		for index, value := range q.Elements {
			if elem > value {
				temp := append([]int{}, q.Elements[index:]...)
				q.Elements = append(q.Elements[:index], elem)
				q.Elements = append(q.Elements, temp...)
				break
			}
		}
		if q.Len == len(q.Elements) {
			q.Elements = append(q.Elements, []int{elem}...)
		}
	}

	q.Len += 1
}

func (q *PriorityQueueAdd) Dequeue() int {
	if q.Len == 0 {
		return 0
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	q.Len -= 1
	return firstElement
}
