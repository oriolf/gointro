package main

import "fmt"

func main() {
	fruta := "fresa"

	switch fruta {
	case "platano":
		fmt.Println("Color amarillo")

	case "naranja":
		fmt.Println("Color naranja")

	case "pera":
		fmt.Println("Color verde")

	case "fresa", "cereza":
		fmt.Println("Color rojo")

	default:
		fmt.Println("Color indefinido")
	}
}
