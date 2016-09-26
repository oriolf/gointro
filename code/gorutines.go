package main

import "fmt"

func showMsg(message string) {
	for i := 0; i < 3; i ++ {	
		fmt.Println(message, i)
	}
}

func main() {
	showMsg("Primero")
	go showMsg("Segundo")
	fmt.Println("Ahora le toca al tercero")
	showMsg("Tercero")
}
