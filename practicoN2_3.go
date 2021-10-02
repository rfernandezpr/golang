package main

import (
	"fmt"
)

const (
	edad int = 10
	sexo     = "Masculino"
)

func main() {

	fmt.Printf("%T\t%v\n", edad, edad)
	fmt.Printf("%T\t%v\n", sexo, sexo)
}
