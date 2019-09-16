package main

import "fmt"

type Object interface {
}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

func (this *List) Length() int {
	cur := this.headNode
	count := 0

	for cur != nil {
		count++
		cur = cur.Next
	}

	return count
}

func (this *List) AddHead(data Object) {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
}

func (this *List) AddTail(data Object) {
	node := &Node{Data: data, Next: nil}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	if index < 0 || index >= this.Length() {
		return
	} else if index == 0 {
		this.headNode = pre.Next
	} else {
		for index != 1 && pre.Next != nil {
			index--
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (this *List) ForeachList() {
	if this.IsEmpty() != true {
		cur := this.headNode
		for cur != nil {
			fmt.Printf("\t%v", cur.Data)
			cur = cur.Next
		}
	}
	fmt.Println()
}

func main() {
	list := List{}
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)
	list.AddTail('a')
	list.AddTail("bcd")
	list.ForeachList()

	list.RemoveAtIndex(0)
	list.ForeachList()
	list.RemoveAtIndex(10)
	list.RemoveAtIndex(-10)
	list.ForeachList()
	list.RemoveAtIndex(1)
	list.ForeachList()
	list.RemoveAtIndex(1)
	list.ForeachList()
}
