package main

import (
	"fmt"
	"math"
)

// Redundar en codigo -FUNCIONES Y FUNCIONES ANONIMAS

// 1. se utiliza la palabra reservada "func"
func nombreFuncion(nombreVariable string) {
	fmt.Println(nombreVariable)
}

// 2. en ocaciones se necesitan multiples argumentos y tipos de datos
// Se puede agrupár los mismos tipos de datos ejemplo: a, b int
func multipleArgumento(a, b int, c float64, d string, e bool) {
	fmt.Println(a, b, c, d, e)
}

// 3. Funcion que retorna un valor.
func retornaValor(base, altura uint) uint {
	return base * altura
}

// 4. Funcion que retoran varios valores
func retoroMultiple(radio float64) (circunferenciaCirculo, areaCirculo float64) {
	return 2 * math.Pi * radio, math.Pi * math.Pow(radio, 2)

}

func main() {
	// se invoca la funcion llamandola por su nombre y escribiendo en los parentesis el valor de los argumentos o variables

	// 1. funcion nombreFuncion
	nombreFuncion("Hola Mundo")

	// 2. funcion multipleArgumneto
	multipleArgumento(1, 2, 3.1415, "hola", false)

	// 3. funcion retornaValor
	areaRectangulo := retornaValor(2, 4)
	fmt.Println("Area Rectangulo: ", areaRectangulo)

	// 4. funcion retornoMultiple
	circunferencia, areaCirculo := retoroMultiple(2)
	fmt.Printf("Circunferencia: %g \nArea de un circulo: %g\n", circunferencia, areaCirculo)

	// 4.1 funcion retornoMultiple, CON UN SOLO RETURN
	_, areaCirculo2 := retoroMultiple(3)
	// SE REEMPLAZA LA VARIABLE A RETORNAR POR UN guion piso "_"
	fmt.Printf("Area de un circulo: %g\n", areaCirculo2)
}
