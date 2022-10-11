package main

import (
	"fmt"
)

// **** LLAVE - VALOR  :: Ejemplo: DICCIONARIO
// 		En go, se define con un la palabra reservada "map"
//		(map[Tipo de variable de la LLAVE]Tipo de variable del VALOR)

func main() {
	urna := make(map[string]int)
	urna["ALCALDE1"] = 10
	urna["ALCALDE2"] = 51
	urna["ALCALDE3"] = 37

	fmt.Println(urna)

	// RECORRIDO DE map

	for k, v := range urna {
		fmt.Println(k, v)
	}

	// IMPRIMIR SOLO LLAVE DE map
	for k := range urna {
		fmt.Println(k)
	}

	// IMPRIMIR SOLO VALOR DE map
	for _, v := range urna {
		fmt.Println(v)
	}

	// PARA VERIFICAR CONTENIDO DE UNA LLAVE =>>
	// SE DEFINE ASI: "variable 1, variable 2 = map[key]"
	// captura LOS SIGUIENTES DATOS DEL map[key} =>>
	// llave = value, verificacion(true, flase) = ok
	value, ok := urna["ALCALDE1"]
	fmt.Println(value, ok)

	// SI LLAVE NO EXISTE "value = 0" & "ok = false"
	v, o := urna["ALCALDE4"]
	fmt.Println(v, o)

	x := urna["ALCALDE3"]
	fmt.Println(x)

	_, y := urna["ALCALDE3"]
	fmt.Println(y)

}
