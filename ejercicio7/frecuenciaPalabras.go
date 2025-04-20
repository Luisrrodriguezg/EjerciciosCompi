package main

import (
	"fmt"
	"strings"
	"unicode"
)

func frecuenciaPalabras(parrafo string) {
	// Separar palabras usando cualquier carácter que NO sea letra
	palabrasSeparadas := strings.FieldsFunc(parrafo, func(r rune) bool {
		// true ⇒ separamos aquí (coma, punto, espacio, número, etc.)
		return !unicode.IsLetter(r)
	})

	// Utilizando map (dict en go) inicializo a las palabras = 0
	resultado := make(map[string]int)
	for _, palabra := range palabrasSeparadas {
		// 1) normalizo a minúsculas
		p := strings.ToLower(palabra)
		// 2) elimino todo lo que no sea letra
		p = strings.TrimFunc(p, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		// 3) si queda vacío, lo salto
		if p == "" {
			continue
		}
		resultado[p] = 0
	}

	// volver a correr para sumar el número de veces que aparece
	for _, palabra := range palabrasSeparadas {
		// **Misma normalización que arriba** antes de contar
		p := strings.ToLower(palabra)
		p = strings.TrimFunc(p, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		if p == "" {
			continue
		}
		resultado[p] = resultado[p] + 1
	}

	// Por último imprimir
	fmt.Println(resultado)
}

func main() {
	frecuenciaPalabras("Hola mundo. Hola clase, hola estudiantes!")
}
