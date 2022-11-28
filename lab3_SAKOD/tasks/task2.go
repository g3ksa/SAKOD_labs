package tasks

import (
	"fmt"
	_ "fmt"
	"lab3_SAKOD/LinkedList"
	"time"
)

func Task2() {
	count := 0
	index := 0
	countf := 0
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	linkedList := LinkedList.LinkedList{}
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	linkedList.PushBack(5)
	linkedList.PushBack(6)
	linkedList.PushBack(7)
	linkedList.PushBack(8)
	linkedList.PushBack(9)
	linkedList.PushBack(10)

	fmt.Println("Реализация на массиве запущена:")
	start := time.Now().Nanosecond()

	for countf < len(array)-1 {
		if array[index] != -1 {
			count++
			if count == 3 {
				array[index] = -1
				if index == len(array)-1 {
					index = 0
				}
				countf++
				count = 0
				index++
			} else {
				if index == len(array)-1 {
					index = 0
				} else {
					index++
				}
			}
		} else {
			if index == len(array)-1 {
				index = 0
			} else {
				index++
			}
		}
	}

	end := time.Now().Nanosecond()
	fmt.Println("%d; %d; %d", start, end, end-start)

	for i := 0; i < len(array); i++ {
		if array[i] != -1 {
			fmt.Println("Уцелевший:", array[i])
		}
	}

	fmt.Println("Реализация на связанном списке запущена:")
	node := linkedList.First
	//start = time.Now()
	for linkedList.GetLength() != 1 {
		count++
		if count == 3 {
			var next = node.Next

			if next == nil {
				next = linkedList.First
			}

			linkedList.Remove(node)
			node = next
			count = 0
		} else {
			if node.Next == nil {
				node = linkedList.First
			} else {
				node = node.Next
			}
		}
	}

	//end = time.Since(start).Nanoseconds()
	fmt.Println("Прошло", end)

	fmt.Print("Уцелевший:")
	linkedList.Print()
}
