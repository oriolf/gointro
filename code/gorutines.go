package main

import "fmt"
import "time"

func showMsg(message string) {
	for i := 0; i < 3; i ++ {	
		fmt.Println(message, i)
		time.Sleep(10*time.Millisecond)
	}
}

func main() {
	showMsg("Primero")
	go showMsg("Segundo")
	fmt.Println("Ahora le toca al tercero")
	showMsg("Tercero")
}
