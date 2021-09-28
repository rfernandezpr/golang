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

	//
	fmt.Println("")

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U\n", s1[i])
	}

	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("En el indice %d el valor es %v\n", i, v)
	}

	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("En el indice %d el valor es %#x\n", i, v)
	}

	fmt.Println("")

	fmt.Printf("Con el vervo %q indico que se imprima el string %s", "%s", s1)

}
