package main

import (
	"fmt"
)

func main() {
	x:= 5
	switch {
	case x == 0: 
		fmt.Printf("El numero es %v", x)
	case x == 1:
		fmt.Printf("El numero es %v", x)
	case x == 2:
		fmt.Printf("El numero es %v", x)
	default:
		fmt.Println("el numero no esta entre 0 y 2")
	}
}