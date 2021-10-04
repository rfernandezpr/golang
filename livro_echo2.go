// Echo2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// usa o identificador vazio para atribuir o indice do for ja que não precisaremos dele
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
