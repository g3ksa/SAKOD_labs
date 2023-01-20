package LinkedList

import (
	"errors"
	"fmt"
	"reflect"
)

func min(args ...int) int {
	min := args[0]
	for _, value := range args {
		if value < min {
			min = value
		}
	}

	return min
}

type Node struct {
	Prev  *Node
	Next  *Node
	Value interface{}
}

type LinkedList struct {
	First  *Node
	Last   *Node
	Length int
}

func (L *LinkedList) IsEmpty() bool {
	return L.First == nil
}

func (L *LinkedList) PushBack(value interface{}) {
	L.Length++
	var p *Node = &Node{
		Next:  nil,
		Prev:  nil,
		Value: value,
	}

	if L.IsEmpty() {
		L.First = p
		L.Last = p
		return
	}

	temp := L.Last
	L.Last.Next = p
	L.Last.Next.Prev = temp
	L.Last = p
}

func (L *LinkedList) Print() {
	item := L.First

	for item != nil {
		if item == L.Last {
			fmt.Printf(" %+v", item.Value)
		} else if item == L.First {
			fmt.Printf("%+v ->", item.Value)
		} else {
			fmt.Printf(" %+v ->", item.Value)
		}
		item = item.Next
	}
	fmt.Println()
}

func (L *LinkedList) GetLength() int {
	return L.Length
}

func (L *LinkedList) Reverse() {
	temp := L.First
	var prev *Node
	L.Last = temp

	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}

	L.First = prev
}

func (L *LinkedList) GetMin() (interface{}, error) {
	if L.IsEmpty() {
		return 0, errors.New("empty list")
	}
	if reflect.TypeOf(L.First.Value).Name() == "string" || reflect.TypeOf(L.First.Value).Name() == "bool" {
		return 0, errors.New("not a numeric value")
	}

	var item = L.First
	var minInt, ok = L.First.Value.(int)
	if !ok {
		var minFloat = L.First.Value.(float64)
		var tempFloat float64
		for item != nil {
			tempFloat = item.Value.(float64)
			if tempFloat < minFloat {
				minFloat = tempFloat
			}
			item = item.Next
		}
		return minFloat, nil
	}
	var tempInt int

	for item != nil {
		tempInt = item.Value.(int)
		if tempInt < minInt {
			minInt = tempInt
		}
		item = item.Next
	}

	return minInt, nil
}

func (L *LinkedList) DeleteAfter(n *Node) error {
	if n == L.Last {
		return errors.New("Невозможно удалить элемент после последнего")
	}

	if n.Next.Next != nil {
		n.Next = n.Next.Next
		n.Next.Prev = n
	} else {
		n.Next = nil
		L.Last = n
	}
	L.Length--
	return nil
}

func (L *LinkedList) Remove(n *Node) error {
	if n == nil {
		return errors.New("Заданного элемента не существует")
	}

	if n.Prev == nil {
		L.First = n.Next
		L.First.Prev = nil
	} else if n.Next != nil {
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	} else {
		n.Prev.Next = nil
		L.Last = n.Prev
	}
	L.Length--
	return nil
}

func (L *LinkedList) Find(value interface{}) (*Node, error) {
	item := L.First

	for item != nil {
		if item.Value == value {
			return item, nil
		}
		item = item.Next
	}
	return nil, errors.New("Элемент не найден")
}

func Repeat(initialList *LinkedList, repeatCnt int) (*LinkedList, error) {
	resultList := LinkedList{}
	if initialList.IsEmpty() {
		return nil, errors.New("Кажется тут нечего повторять(пустой список)")
	}
	for i := 0; i < repeatCnt; i++ {
		item := initialList.First
		for item != nil {
			resultList.PushBack(item.Value)
			item = item.Next
		}
	}
	return &resultList, nil
}
