package main

import (
	"fmt"
)

func main() {
	x := "Pinguino"

	if x == "Batman" {
		fmt.Println(x)
	} else if x == "Robin" {
		fmt.Printf("No es Batman es %v\n", x)
	} else {
		fmt.Printf("No es Batman ni es Robin es %v\n", x)
	}
	
}