package main

import (
	"fmt"
)

func main() {
	for x:= 10; x <= 100; x++ {
		fmt.Printf("cunado dividimos %v entre 4, el resto es %v\n", x, x%4)
	}
}