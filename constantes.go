package main

import (
	"fmt"
)

const a = 42
const b = 42.32
const c = "Eduar Tua"
const d string = "Roger Fernandez"

//las constantes pueden ser con tipo o sin tipo, son mas flexibles sin tipo

type nombre string

var e nombre

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	e = c

	//e = d
	// si hacemos e = d genera un error de tipo porque string no es igual a nombre

	fmt.Println(e)

}
