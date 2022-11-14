package main

import (
	"fmt"
)

func main() {
	tree := BinaryTree{}
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(11)
	tree.Insert(33)
	tree.Insert(89)

	//fmt.Println(tree.Search(50, nil))
	//fmt.Println(tree.Search(89, nil))
	//fmt.Println(tree.Search(11, nil))
	//fmt.Println(tree.Search(1, nil))
	//fmt.Println(tree.Search(2, nil))
	//fmt.Println(tree.Search(3, nil))

	//fmt.Println(tree.Delete(50))
	//fmt.Println(tree.Search(50, nil))

	if tree.Root != nil {
		tree.Root.printSort()
	}
}

type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) Search(target int, parent *Node) (*Node, bool) { // 삭제를 위해 target 부모를 리턴한다.
	if b.Root == nil {
		return nil, false
	}

	if b.Root.Data == target {
		return parent, true
	}

	if b.Root.Left != nil {
		leftBinary := BinaryTree{Root: b.Root.Left}
		if parent, isExist := leftBinary.Search(target, b.Root); isExist {
			return parent, true
		}
	}

	if b.Root.Right != nil {
		rightBinary := BinaryTree{Root: b.Root.Right}
		if parent, isExist := rightBinary.Search(target, b.Root); isExist {
			return parent, true
		}
	}
	return nil, false
}

func (b *BinaryTree) Insert(newVal int) error {
	insert := &Node{Data: newVal}
	if b.Root == nil {
		b.Root = insert
		return nil
	}

	find := b.Root
	for {
		fmt.Print(find.Data, "-> ")
		if find.Data > newVal {
			if find.Left == nil {
				find.Left = &Node{Data: newVal}
				fmt.Println("left", newVal)
				return nil
			}
			find = find.Left
			continue
		} else if find.Data < newVal {
			if find.Right == nil {
				find.Right = &Node{Data: newVal}
				fmt.Println("right", newVal)
				return nil
			}
			find = find.Right
		}
	}

	return nil
}

func (b *BinaryTree) Delete(target int) error {
	//  삭제 타겟의 부모를 찾는다.
	parent, isExist := b.Search(target, nil)
	if !isExist {
		return fmt.Errorf("삭제 대상이 없음")
	}
	del := &Node{} // 삭제 대상
	child := ""
	if parent == nil {
		del = b.Root
	} else if parent.Right != nil && parent.Right.Data == target {
		del = parent.Right
		child = "Right"
	} else if parent.Left != nil && parent.Left.Data == target {
		del = parent.Left
		child = "Left"

	}

	if del.Left == nil && del.Right == nil { //	자식이 없으면 그냥 삭제한다.
		if parent == nil { // root 만 있는 트리인 경우 root 삭제
			b.Root = nil
			return nil
		}
		parent.DeleteChild(child)
		return nil
	}
	if del.Left == nil { //	자식이 하나면 자식을 삭제된 노드 위치에 넣는다. right 만 있을 때
		*del = *del.Right
		return nil
	}
	if del.Right == nil { // 자식이 하나면 자식을 삭제된 노드 위치에 넣는다. left 만 있을 때
		*del = *del.Left
		return nil
	}

	//	자식이 둘이면 후속자 노드(삭제하고자 하는 값보다 바로 다음 큰 값) 를 찾아서 값을 삭제할 노드 위치에 복사하고, 후속자 노드는 삭제한다.
	//  후속자 노드 찾기
	//	삭제할 노드의 right 자식에서
	nextChild := "Right"
	nextParent := del
	next := del.Right
	//	left 자식이 있을 때 까지 left 로 이동, left 자식이 없는 노드가 후속자 노드
	for {
		if next.Left == nil {
			break
		}
		nextParent = next
		next = next.Left
		nextChild = "Left"
	}
	//	삭제 타겟자리에 후속자 노드 값을 넣는다.
	del.Data = next.Data
	//	후속노드에 right 자식이 있는 경우, right 자식을 후속자 노드 부모의 왼쪽에 넣는다.
	if next.Right != nil {
		*next = *next.Right
	}
	// 후속자 노드는 삭제한다.
	nextParent.DeleteChild(nextChild)
	return nil
}

func (n *Node) DeleteChild(child string) {
	if n == nil {
		return
	}
	if child == "Right" {
		n.Right = nil
	}
	if child == "Left" {
		n.Left = nil
	}
}

func (n *Node) printSort() {
	if n.Left != nil {
		n.Left.printSort()
	}
	fmt.Println("tree 데이터를 작은 순으로 출력...", n.Data)
	if n.Right != nil {
		n.Right.printSort()
	}
}
