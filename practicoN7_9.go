package main

import (
	"fmt"
)

func main() {

	deporteFav := ""

	switch deporteFav {
	case "futboll":
		fmt.Println("Su deporte favorito es el %v", deporteFav)
	case "basketboll":
		fmt.Println("Su deporte favorito es el %v", deporteFav)
	case "baseball":
		fmt.Println("Su deporte favorito es el %v", deporteFav)
	default:
		fmt.Println("No tiene deporte favorito")
	}
}