package main

import "fmt"

func main() {
	a := 2
	var c uint = uint(a)
	
	fmt.Printf("a is of type %T, value: %v\n", a, a)
	fmt.Printf("c is of type %T, value: %v\n", c, c)
}
