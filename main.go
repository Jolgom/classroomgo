package main

import "fmt"

// *** 	CONDICIONALES

// SWITCH : Utilizar multiples condiciones anidadas
func main() {

	// con condicion especificada
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es Impar")
	}

	// sin condicion especificada
	valor := 50
	switch {
	case valor > 100:
		fmt.Println("Valor Es mayor que 100")
	case valor < 0:
		fmt.Println("Valor Es menor que 0")
	default:
		fmt.Println("No es una condicion")
	}
}
