package main

import (
	"fmt"
)

func main() {

	s1 := "Hola mundo"

	//row string literal
	s2 := `Esta es una linea 
	partida.`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es: %T\n", s2)

	fmt.Println("")

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T", bs)

}
