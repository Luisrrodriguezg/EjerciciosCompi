package main

import "strings"

func rotarPalabrasK(palabras string, k int) string {
	var resultado string
	//Separar palabras por espacios " "
	palabrasSeparadas := strings.Split(palabras, " ")
	//Separacion
	largo := len(palabrasSeparadas) - k
	//Dos arreglos que se les repartiran 'palabrasSeparadas' dividido en la posicion k
	principioArreglo := palabrasSeparadas[0:largo]
	finalArreglo := palabrasSeparadas[largo : largo+k]
	//Añadir todas al final
	for _, word := range principioArreglo {
		finalArreglo = append(finalArreglo, word)
	}
	//Convertirlas en String
	resultado = strings.Join(finalArreglo, " ")
	return resultado
}
func main() {
	print(rotarPalabrasK("uno dos tres cuatro", 1))
}
