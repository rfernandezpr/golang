package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	// o mesmo que
	// a = iota
	// b = iota
	// c = iota
)

// cada vez que aparece la palabra clave const entonces iota se reinicia a 0
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
