package main

import (
	"fmt"
)

/*
{1, 2, 3, 4, 5, 6, 7, 8, 9}
		   1
		  / \
		 2   3
		/ \ / \
	   4  5 6  7
	  / \
	 8   9
*/
/*
{9, 4, 5, 11, 28, 1, 3, 2}
		   9
		  / \
		 4  11
		/ \   \
	   1   5  28
		\
		 3
        /
       2
*/

type BTNode struct {
	value int
	left, right *BTNode
}

var(
	tree []BTNode
)

// create complete binary tree (linked list)
func createBTree() *BTNode {
	treeArray := []int{9, 4, 5, 11, 28, 1, 3, 2}

	var root *BTNode
	for i := 0; i < len(treeArray); i++ {
		root = insertBTNode(root, treeArray[i])
	}

	fmt.Print("先序遍历（递归）：")
	preOrder(root)
	fmt.Println()

	fmt.Print("中序遍历（递归）：")
	middleOrder(root)
	fmt.Println()

	fmt.Print("后序遍历（递归）：")
	postOrder(root)
	fmt.Println()

	fmt.Print("先序遍历（迭代）：")
	preOrderReverse(root)
	fmt.Println()

	return root
}

func insertBTNode(root *BTNode, m int) *BTNode {
	// add one by one to the last in order (complete binary tree)
	node := BTNode{m, nil, nil}
	fmt.Printf("%d ", m)

	var current *BTNode = root
	var last *BTNode

	if root == nil {
		root = &node
	} else {
		for current != nil {
			last = current	// 找到新节点要挂的父节点
			if node.value > current.value {
				current = current.right
			} else if node.value < current.value {
				current = current.left
			} else {
				// 相同值的node不再加入
			}
		}

		if node.value > last.value {
			last.right = &node
			fmt.Printf("parent is %d , it is right child.", last.value)
		} else if node.value < last.value {
			last.left = &node
			fmt.Printf("parent is %d , it is left child.", last.value)
		} else {
			// 相同的节点不再处理
		}
	}
	fmt.Println()
	return root
}

// recursive formula
func preOrder(root *BTNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.value)
	preOrder(root.left)
	preOrder(root.right)
}
func middleOrder(root *BTNode) {
	if root == nil {
		return
	}
	middleOrder(root.left)
	fmt.Printf("%d ", root.value)
	middleOrder(root.right)
}
func postOrder(root *BTNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.value)
}
// iterative method
func preOrderReverse(root *BTNode) {
	if root == nil {
		fmt.Println("null binary tree!")
		return
	}
	var current *BTNode
	var stack []*BTNode
	stack = append(stack, root)
	i := 0

	for current = stack[i]; current != nil; current = stack[i] {
		fmt.Printf("%s ", current.value)
		i = i - 1
		stack := stack[:i]

		for current.left != nil {
			stack = append(stack, current.left)
			i++
			current = current.left
			fmt.Printf("%s ", current.value)
		}
		if current.right != nil {
			stack = append(stack, current.right)
			i++
			current = current.right
			fmt.Printf("%s ", current.value)
		}
	}
}


// create complete binary tree (array or slice)
func createBTreeArray(n int) []BTNode {
	tree = make([]BTNode, n)
	tree[0].value = 1
	tree[0].left = nil
	tree[0].right = nil

	for i := 1; i < n; i++ {
		num := i + 1
		tree[i].value = num
		tree[i].left = nil
		tree[i].right = nil

		parent := num / 2 - 1
		if (0 == num % 2) {
			// i is left child
			tree[parent].left = &tree[i]
		} else {
			// i is left child
			tree[parent].right = &tree[i]
		}
	}

	//  displayTree(tree)

	return tree
}
// display one by one in order
func displayTree(tree []BTNode) {
	if tree == nil {
		fmt.Println("the tree is empty")
	}
	for i := 0; i < len(tree); i ++ {
		fmt.Print(tree[i].value)
		if tree[i].left != nil {
			fmt.Print(" left is ", tree[i].left.value)
		}
		if tree[i].right != nil {
			fmt.Print(" right is ", tree[i].right.value)
		}
		fmt.Println()
	}
}

func main() {
	createBTree()
}