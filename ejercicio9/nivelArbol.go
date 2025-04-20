package main

import "fmt"

// mismo arbol del ejercicio anterior
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nivelesArbol(root *TreeNode) [][]int {
	//resultado como slice de slices
	var resultado [][]int
	if root == nil {
		return resultado
	}

	cola := []*TreeNode{root} // inicio con la raíz
	for len(cola) > 0 {
		nivelSize := len(cola)
		//cro el arreglo del tamaño de los hijos del arbol
		nivel := make([]int, 0, nivelSize)

		// procesar el nivel actual
		for i := 0; i < nivelSize; i++ {
			nodo := cola[0]
			cola = cola[1:]                 // desencolar
			nivel = append(nivel, nodo.Val) // guarda el valor
			// encolar hijos
			if nodo.Left != nil {
				cola = append(cola, nodo.Left)
			}
			if nodo.Right != nil {
				cola = append(cola, nodo.Right)
			}
		}

		resultado = append(resultado, nivel) // añade este nivel al resultado
	}

	return resultado
}

func main() {
	root := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil},
	}
	fmt.Println("Niveles del árbol:", nivelesArbol(root)) // Esperado: [[1], [2,3], [4,5]]
}
