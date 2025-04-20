package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Palabras para analizar
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//anagramas organizados para comparar
	var sortedWords []string
	//Resultado final
	var result [][]string

	//For para mandar las palabras a analizar
	for _, currentWord := range words {
		//Division Palabras
		dividedWords := strings.Split(currentWord, "")
		//Organizar palabras en orden alfabético
		sort.Strings(dividedWords)
		//Convertir en string
		anagrama := strings.Join(dividedWords, "")
		//Flag para añadir luego al la lista de anagrama y contador para añadir al resulatado
		sortedWordFlag := false
		contador := 0
		//For para verificar si hace parte a una familia de anagrama para que sea añadida a result
		for _, currentOrganizedW := range sortedWords {
			if anagrama == currentOrganizedW {
				sortedWordFlag = true
				result[contador] = append(result[contador], currentWord)
				break //Si existe rompe y añade
			}
			contador++
		} //De lo contrario, añada a sortedwords y la última fila de resultado
		if sortedWordFlag == false {
			sortedWords = append(sortedWords, anagrama)
			//Se añade una lista con un elemento
			result = append(result, []string{currentWord})
		}
	}
	//Mostrar resultado
	for i, grupo := range result {
		fmt.Printf("Grupo %d: %v\n", i+1, grupo)
	}

}
