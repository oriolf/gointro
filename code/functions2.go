package main

import "fmt"

func swapName(name string, surname string) (string, string) {
	return surname, name
}

func main() {
	first, second := swapName("Terry", "Colby")
	fmt.Println("I am", first, second )
}
