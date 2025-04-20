package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func esSimetrico(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return sonEspejo(root.Left, root.Right)
}

func sonEspejo(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val &&
		sonEspejo(a.Left, b.Right) &&
		sonEspejo(a.Right, b.Left)
}

func main() {

	root := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}},
	}
	fmt.Println("Árbol simétrico:", esSimetrico(root))
}
