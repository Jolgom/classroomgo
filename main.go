package main

import "fmt"

// *** 	KEYWORDS

// break
// continue

// defer: ejecuta line de codigo al final
//
//	la buena practica es usar solo defer por codigo
func main() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue y break	se utiliza en un for

	for i := 0; i < 10; i++ {
		fmt.Print(i)

		if i == 2 {
			fmt.Print(" Es 2 ")
			continue
		}
		if i == 8 {
			fmt.Print(" Fin Forzado ")
			break
		}
	}

}
