//Demuestra el uso de los alias para los tipos de variables y el uso del paquete runtime
package main

import (
	"fmt"
	"runtime"
)

//byte es un alias para el tipo uibt8 que acepta solo numeros del 0 al 255 solo positivos
var x byte

//rune es un alias para el tipo int32
var y rune

func main() {

	x = 254
	y = 255

	fmt.Println("El valor de x es: ", x)
	fmt.Printf("El tipo de x es: %T\n", x)

	fmt.Println("El valor de y es: ", y)
	fmt.Printf("El tipo de y es: %T\n\n", y)

	fmt.Println("El Sistema Operativo es: ", runtime.GOOS)
	fmt.Println("La arquitectura usada es:", runtime.GOARCH)
}
