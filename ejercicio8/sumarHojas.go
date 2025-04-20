package main

import "fmt"

// Definición básica de nodo de árbol binario
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumarHojas(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// Si es hoja, devuelvo su valor
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	// Sumo recursivamente las hojas de subárbol izquierdo y derecho
	return sumarHojas(root.Left) + sumarHojas(root.Right)
}

func main() {
	root := &TreeNode{
		5,
		&TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{8, nil, &TreeNode{10, nil, nil}},
	}
	fmt.Println("Suma de hojas:", sumarHojas(root)) // Esperado: 15
}
