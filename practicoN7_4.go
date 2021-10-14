package main

import (
	"fmt"
)

func main() {
	i:=1971;
	for {
		if(i > 2021) {
			break
		} 
		fmt.Println(i)
		i++
	}
}