package main

import "fmt"

func formalName(name string) string {
	return "Mr. " + name
}

func main() {
	fmt.Println("Welcome " + formalName("Terry"))
}
