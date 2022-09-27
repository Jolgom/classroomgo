package main

import "fmt"

// *** CICLOS
// EN go SOLO EXISTE UN CICLO, << for >>

func main() {
	// 01.	for basico: condicional (FOR TRADICIONAL)
	// 		IDNICAR AL INICIO PARA QUE ITERE, INCREMENTE Y FINALICE SEGUN LA CONDICION.
	for i := 0; i <= 3; i++ {
		fmt.Println("la iternacia es: ", i)
	}

	// 02.	for while
	// 		SE DEBE CUMPLIR UNA CONDICION PARA QUE FINALICE EL CICLO.
	//		la condicion se define previamente
	contador := 4
	for contador < 8 {
		fmt.Println("la iternacia es: ", contador)
		contador++
	}

	// 04. for REVERSA
	contador = 11
	for 8 <= contador {
		fmt.Println("la iternacia es: ", contador)
		contador--
	}
}
