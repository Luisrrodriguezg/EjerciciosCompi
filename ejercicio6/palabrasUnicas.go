package main

import (
	"fmt"
	"sort"
	"strings"
)

func palabraUnicas(palabras []string) []string { //Reutilizando el código de anagramas
	var resultado []string
	for _, word := range palabras {
		//Division Palabras
		dividedWords := strings.Split(word, "")
		//Organizar palabras en orden alfabético
		sort.Strings(dividedWords)
		//con las letras organizadas, averiguo si se repiten letras
		flagRepetidas := false
		for i := 0; i < len(dividedWords)-1; i++ {
			if dividedWords[i] == dividedWords[i+1] {
				flagRepetidas = true
				break
			}
		}
		//en caso de que la bandera siga siendo falsa, agregar
		if !flagRepetidas {
			resultado = append(resultado, word)
		}
	}
	return resultado
}

func main() {
	prueba := []string{"azul", "ojo", "verde", "a"}
	prueba2 := []string{"hola", "casa", "perro", "luz"}
	fmt.Println(palabraUnicas(prueba))
	fmt.Println(palabraUnicas(prueba2))
}
