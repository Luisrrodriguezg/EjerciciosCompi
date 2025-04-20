package main

import (
	"fmt"
)

func eliminarDuplicadoLista(lista []string) []string {
	//Utilizando map (dict en go) adjuntamos indices a las palabras
	m := make(map[string]int)
	for index, palabra := range lista {
		m[palabra] = index //si las palabras se repiten, las llaves del mismo nombre tendran el valor mayor
	}
	var resultado []string
	//recorre nuevamente para añadir los últimos indices en el orden correcto
	for index, palabra := range lista {
		if m[palabra] == index {
			resultado = append(resultado, palabra)
		}
	}
	return resultado
}
func main() { //el método es sensible a mayusculas
	lista := []string{"apple", "banana", "pear", "melon", "pear", "banana", "melon"}
	resultado := eliminarDuplicadoLista(lista)
	fmt.Print(resultado)
}
