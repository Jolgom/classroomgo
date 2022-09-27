package main

import (
	"fmt"
	"log"
	"strconv"
)

// *** 	CONDICIONALES

// if
func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// compuerta AND
	if valor1 == 0 && valor2 == 2 {
		fmt.Println("ambas Son verdad")
	} else {
		fmt.Println("AL MENOS UNA ES FALSA")
	}

	// compuerta OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("AL MENOS UNA ES VERDADERA")
	}
	// CONVERTIR TEXTO A NUMERO
	value, err := strconv.Atoi("20")
	if err != nil { // nil : palabra reservada del lenguaje para indicar que no hay error
		log.Fatal(err)
	}
	fmt.Println("Valor: ", value)
}
