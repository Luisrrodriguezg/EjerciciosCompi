package main

import (
	"fmt"
)

func listasIntercaladas(lista1 []string, lista2 []string) []string {
	resultado := []string{}
	//Comparar el tamaño mayor para intercalarlas, el resto copiar array.
	largo1 := len(lista1)
	largo2 := len(lista2)
	//añadido intercalado por condicion
	if largo1 >= largo2 {
		for i, v := range lista2 {
			resultado = append(resultado, lista1[i])
			resultado = append(resultado, v)
		}
		//Añadir el resto
		resultado = append(resultado, lista1[largo2:largo1]...)
	} else {
		for i, v := range lista1 {
			resultado = append(resultado, lista2[i])
			resultado = append(resultado, v)
		} //Añadir el resto
		resultado = append(resultado, lista2[largo1:largo2]...)
	}
	return resultado
}
func main() {
	prueba := listasIntercaladas([]string{"a", "v", "f", "gf"}, []string{"1", "3"})
	fmt.Print(prueba)
}
