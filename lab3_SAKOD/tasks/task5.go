package tasks

import (
	"fmt"
	queue "lab3_SAKOD/Queue"
)

func Task5() {
	// priorityQueue := queue.PriorityQueueAdd{}

	// priorityQueue.Enqueue(4)
	// fmt.Println(priorityQueue.Elements)

	// priorityQueue.Enqueue(6)
	// priorityQueue.Enqueue(5)
	// priorityQueue.Enqueue(3)
	// priorityQueue.Enqueue(2)
	// priorityQueue.Enqueue(1)
	// priorityQueue.Dequeue()
	// fmt.Println(priorityQueue.Elements)

	priorityQueue := queue.PriorityQueueDelete{}

	priorityQueue.Enqueue(4)
	priorityQueue.Enqueue(7)
	priorityQueue.Enqueue(1)
	priorityQueue.Enqueue(3)
	priorityQueue.Enqueue(9)
	fmt.Println(priorityQueue.Elements)

	elem := priorityQueue.Dequeue()

	fmt.Println(priorityQueue.Elements, elem)

	elem = priorityQueue.Dequeue()

	fmt.Println(priorityQueue.Elements, elem)

}
