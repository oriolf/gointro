package main

import "fmt"

func add(firstValue, secondValue int) int{
	return firstValue + secondValue
}

func main() {
	var a, b = 2, 4
	if c := add(a, b); c < 10 {
		fmt.Println("No llega a diez")
	} else {
		fmt.Println("Mayor que diez")
	}
}
