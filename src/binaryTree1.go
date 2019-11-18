package main

import "reflect"

type BTNode struct {
	data  int
	left  BTNode
	right BTNode
}

func (root BTNode) BSearch() {
	var queue []BTNode
	queue = append(queue, root)
	for i := 0; !queue[i].IsEmpty(); i++ {
		node := queue[i]
		if !node.left.IsEmpty() {
			queue = append(queue, node.left)
		}
		if !node.right.IsEmpty() {
			queue = append(queue, node.right)
		}
	}
}

func (node BTNode) IsEmpty() bool {
	return reflect.DeepEqual(node, BTNode{})
}
