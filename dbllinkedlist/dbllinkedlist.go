package dbllinkedlist

import "fmt"

type Node struct {
	Prev  *Node
	Next  *Node
	Key   string
	Value interface{}
}

type DblLinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *DblLinkedList) DeleteListHead() {
	if dll.Head != nil && dll.Head.Next != nil {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	} else if dll.Head != nil {
		dll.Head = dll.Head.Next
		dll.Tail = nil
	}
}

func (dll *DblLinkedList) SetListHead(n *Node) {
	if dll.Head == nil {
		dll.Head = n
		dll.Tail = n
	} else {
		prevHead := dll.Head
		newHead := n
		newHead.Next = prevHead
		prevHead.Prev = newHead
		dll.Head = newHead
	}
}

func (dll *DblLinkedList) SetListTail(n *Node) {
	if dll.Head == nil {
		dll.Head = n
		dll.Tail = n
	} else {
		prevTail := dll.Tail
		prevTail.Next = n
		n.Prev = prevTail
		dll.Tail = n
	}
}

func (dll *DblLinkedList) DeleteListTail() {
	if dll.Tail != nil && dll.Tail.Prev != nil {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	} else if dll.Tail != nil {
		dll.Tail = dll.Tail.Prev
		dll.Head = nil
	}
}

func (dll *DblLinkedList) DeleteNode(n *Node) {
	if n.Key == dll.Head.Key {
		dll.DeleteListHead()
		return
	}

	if n.Key == dll.Tail.Key {
		dll.DeleteListTail()
		return
	}

	left := n.Prev
	right := n.Next
	left.Next = right
	right.Prev = left
}

func (dll *DblLinkedList) FindAndDeleteNode(key string) {
	if dll.Head != nil {
		ptr := dll.Head
		for ptr != nil {
			if ptr.Key == key {
				// remove the node
				dll.DeleteNode(ptr)
			}
			ptr = ptr.Next

		}
	}

}

func (dll DblLinkedList) String() string {
	if dll.Head == nil {
		return ""
	}
	tmp := dll.Head
	var linkView string
	for tmp != nil {
		linkView = linkView + "-" + fmt.Sprintf("{key: %s, value: %s}", tmp.Key, tmp.Value)
		tmp = tmp.Next
	}
	return linkView
}
