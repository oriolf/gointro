package main

import "fmt"

func main() {
	a := 2
	var c uint = uint(a)
	
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("c is of type %T\n", c)
}
