package main

import (
	"fmt"
	"strconv"
)

func rangoResumido(listaNums []int) []string {
	var result []string
	n := len(listaNums)
	// Recorre mientras quede al menos un elemento
	for i := 0; i < n; {
		inicio := listaNums[i]
		j := i
		// Avanza j mientras los números sean consecutivos
		for j+1 < n && listaNums[j+1] == listaNums[j]+1 {
			j++
		}
		fin := listaNums[j]
		// Escribir "x->y" o solo "x"
		if inicio == fin {
			result = append(result, strconv.Itoa(inicio))
		} else {
			result = append(result,
				fmt.Sprintf("%d->%d", inicio, fin),
			)
		}
		// Saltar al siguiente segmento
		i = j + 1
	}

	return result
}

func main() {
	lista := []int{1, 2, 3, 4, 6, 7, 9, 23}
	resultado := rangoResumido(lista)
	for _, r := range resultado {
		fmt.Println(r)
	}
}
