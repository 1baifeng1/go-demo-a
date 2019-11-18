package main

import "fmt"

type BtNode struct {
	data  int
	left  *BtNode
	right *BtNode
}

func (root *BtNode) BSearch() {
	if root == nil {
		return
	}

	queue := []*BtNode{root}
	for len(queue) != 0 {
		node := queue[0]
		fmt.Print("%d\t", node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
}
