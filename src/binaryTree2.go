package main

import (
	"fmt"
)

type BtNode struct {
	data  int
	left  *BtNode
	right *BtNode
}

func main() {
	n1 := BtNode{1, nil, nil}
	n2 := BtNode{2, nil, nil}
	n3 := BtNode{3, nil, nil}
	n4 := BtNode{4, nil, nil}
	n5 := BtNode{5, nil, nil}
	n6 := BtNode{6, nil, nil}
	n7 := BtNode{7, nil, nil}
	n8 := BtNode{8, nil, nil}
	n9 := BtNode{9, nil, nil}
	n10 := BtNode{10, nil, nil}
	n11 := BtNode{11, nil, nil}
	n1.left = &n2
	n1.right = &n3
	n2.left = &n4
	n3.left = &n5
	n3.right = &n6
	n5.left = &n7
	n6.right = &n8
	n7.left = &n9
	n7.right = &n10
	n10.left = &n11

	fmt.Print("BSearch():")
	n1.BSearch()
	fmt.Println("GetMaxDepth():", n1.GetMaxDepth())
	fmt.Println("GetKDepthNode():", n1.GetKDepthNode(4))
	fmt.Println("GetAllNode():", n1.GetAllNode())
	fmt.Println("GetLeafNode():", n1.GetLeafNode())

	fmt.Println("=====================")
}

// Go中math.Max()是默认比较的float参数
func GetMax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func GetMin(a int, b int) int {
	// Go 不支持三元表达式
	if a > b {
		return b
	}
	return a
}

// 广度优先打印所有节点
func (root *BtNode) BSearch() {
	if root == nil {
		return
	}

	queue := []*BtNode{root}
	for len(queue) != 0 {
		node := queue[0]
		fmt.Printf("%d\t", node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
	fmt.Println()
}

// 求二叉树最大深度
func (root *BtNode) GetMaxDepth() int {
	if root == nil {
		return 0
	}
	left := root.left.GetMaxDepth()
	right := root.right.GetMaxDepth()
	return GetMax(left, right) + 1
}

// 求二叉树中第K层节点个数
func (root *BtNode) GetKDepthNode(k int) int {
	if root == nil || k < 1 {
		return 0
	}

	if k == 1 {
		return 1
	}
	left := root.left.GetKDepthNode(k - 1)
	right := root.right.GetKDepthNode(k - 1)
	return left + right
}

// 求二叉树中节点总个数：左右子树上节点个数加上该节点
func (root *BtNode) GetAllNode() int {
	if root == nil {
		return 0
	}

	left := root.left.GetAllNode()
	right := root.right.GetAllNode()
	return left + right + 1
}

// 求二叉树中叶子节点个数：只有左右子树都是空的节点是叶子节点，不是叶子点点的话就继续查它的子树
func (root *BtNode) GetLeafNode() int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}
	left := root.left.GetLeafNode()
	right := root.right.GetLeafNode()
	return left + right
}
