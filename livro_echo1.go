// Echo1 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d\t%v\n", i, os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("")
	fmt.Println(s)
}
