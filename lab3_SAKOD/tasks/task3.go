package tasks

import (
	"fmt"
	_ "fmt"
	"lab3_SAKOD/LinkedList"
)

func Task3() {
	var array = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	linkedList := LinkedList.LinkedList{}
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	linkedList.PushBack(5)
	linkedList.PushBack(6)
	linkedList.PushBack(7)
	linkedList.PushBack(8)
	linkedList.PushBack(9)
	linkedList.PushBack(10)
	linkedList.PushBack(11)
	linkedList.PushBack(12)
	linkedList.PushBack(13)
	linkedList.PushBack(14)
	linkedList.PushBack(15)
	linkedList.PushBack(16)
	linkedList.PushBack(17)
	linkedList.PushBack(18)
	linkedList.PushBack(19)
	linkedList.PushBack(20)
	linkedList.PushBack(21)

	for j := 2; j < len(array); j++ {
		for i := 0; i < len(array); i++ {
			if array[i]%j == 0 && array[i] != j {
				array[i] = -1
			}
		}
	}

	for i := 0; i < len(array); i++ {
		if array[i] != -1 {
			fmt.Print(array[i], " ")
		}
	}

	fmt.Println()

	for j := 2; j < linkedList.GetLength(); j++ {
		node := linkedList.First
		for node != nil {
			var next = node.Next
			if node.Value.(int)%j == 0 && node.Value != j {
				linkedList.Remove(node)
			}
			node = next
		}
	}

	linkedList.Print()
}
