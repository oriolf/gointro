package main

import "fmt"

func main(){
	count := 0
	for i := 0; i < 10; i++ {
		count += 1
	}
	fmt.Println("Value: ", count)
}
