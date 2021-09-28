package main

import (
	"fmt"
)

const (
	//desecha el valor inicial de iota 0
	_ = iota
	//implementacion de iota con bitshifting
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a)

	//mueve el bit mas significativo a 1 posicion a la izquierda
	b := a << 1
	fmt.Printf("%d\t\t%b", b, b)

	fmt.Println("")

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
