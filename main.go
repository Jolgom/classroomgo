package main

import "fmt"

// *** LISTAS
// 		array - slice

func main() {

	//	Array: Se define y Se idica entre corchetes "[ ]" el tamaño del arreglo
	//		Se indica el tipo de variable (int,float64, string, bool, etc)
	//		ASí Se crea un arreglo definido en tamaño "n" inmutable
	// 		sus valores se pueden modificar. mas no su tamaño
	var listaArray [4]int // Por defecto su valor es "0"
	fmt.Println(listaArray)
	listaArray[0] = 1
	listaArray[1] = 2

	// len: Tamaño - cap: Maxima capacidad
	fmt.Println(listaArray, len(listaArray), cap(listaArray))

	//	slice: Se define similar a un array, PERO NO Se idica el TAMAÑO de arreglos
	//		Se indica el tipo de variable (int,float64, string, bool, etc)
	//		Se crea un arreglo DE tamaño VARIABLE

	listaSlice := []int{0, 8, 9, 3, 4, 5, 6} // Por defecto su valor es "0"
	fmt.Println(listaSlice)
	listaSlice[1] = 1
	listaSlice[2] = 2

	// len: Tamaño - cap: Maxima capacidad
	fmt.Println(listaSlice, len(listaSlice), cap(listaSlice))

	// Métodos en el slice [vorde interno:vorde externo]
	//		Vorde interno es inclusivo. el valor se icluye en la consulta u operacion
	//		Vorde externo es excluyente. no se incluye el resultado.

	fmt.Print(listaSlice[0])
	fmt.Print(listaSlice[:3])
	fmt.Print(listaSlice[2:4])
	fmt.Println(listaSlice[4:])
	//	Los slice son mutables, se pueden agregar elementos
	//			Se agregan elemetos con el metodo append
	listaSlice = append(listaSlice, 7)
	fmt.Println(listaSlice, len(listaSlice), cap(listaSlice))

	//			Se pueden agregar listas dentro de otra lista con metodo append
	otroSlice := []int{7, 8, 9}
	listaSlice = append(listaSlice, otroSlice...)
	fmt.Println(listaSlice, len(listaSlice), cap(listaSlice))

	// RECORRER SLICE
	slice := []string{"Hola", "Como", "Te", "a", "ido"}
	for i := range slice {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	for _, valor := range slice {
		fmt.Print(valor, " ")
	}

}
