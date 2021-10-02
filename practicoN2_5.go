package main

import (
	"fmt"
)

func main() {

	//string literal no interpretado
	endereco := `Rua João Denbinski 2380
cep: 21830-000 
"Bloco A11 Apto 22"
Curitiba-PR Brasil`

	fmt.Println(endereco)
}
