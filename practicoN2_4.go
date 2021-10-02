package main

import (
	"fmt"
)

func main() {
	v := 42
	si := v << 1

	fmt.Printf("%d\t%b\t%#x\n", v, v, v)
	fmt.Printf("%d\t%b\t%#x\n", si, si, si)

}
