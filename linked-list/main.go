package main

import (
	"errors"
	"fmt"
)

func main() {
	linkedList := LinkedList{}
	linkedList.Push(3)
	linkedList.Push(2)
	linkedList.Push(1)

	fmt.Println(linkedList.Get(0)) // 1
	fmt.Println(linkedList.Get(1)) // 2
	fmt.Println(linkedList.Get(2)) // 3

	fmt.Println("find 8:", linkedList.Find(8)) //false
	fmt.Println("find 3:", linkedList.Find(3)) //true
	fmt.Println("find 1:", linkedList.Find(1)) //true

	fmt.Println("delete: ", linkedList.Delete()) // true (delete 1)
	fmt.Println("find 1:", linkedList.Find(1))   //false

	fmt.Println(linkedList.Get(0)) // 2
	fmt.Println(linkedList.Get(1)) // 3
	fmt.Println(linkedList.Get(2)) // error
}

type LinkedList struct {
	First  *Node
	Last   *Node
	Length int
}

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func toNext(node *Node) *Node {
	if node == nil || node.Next == nil {
		return nil
	}
	return node.Next
}

func (list *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.Length {
		return 0, errors.New("index not valid")
	}
	now := list.First
	for i := 0; i < index; i++ {
		now = toNext(now)
	}
	return now.Data, nil
}

func (list *LinkedList) Push(data int) {
	node := Node{nil, data, nil}
	if list.First == nil {
		list.First = &node
		list.Last = &node
		list.Length = 1
	} else {
		node.Next = list.First
		list.First.Prev = &node
		list.First = &node
		list.Length++
	}
}

func (list *LinkedList) Find(target int) bool {
	if list.Length == 0 {
		return false
	}
	now := list.First
	for now != nil {
		if now.Data == target {
			return true
		}
		now = toNext(now)
	}
	return false
}

func (list *LinkedList) Delete() bool {
	if list.Length == 0 {
		return false
	}
	list.First = list.First.Next
	list.Length--
	if list.First == nil {
		list.Last = nil
	}
	return true
}
