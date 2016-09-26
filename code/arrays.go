package main

import "fmt"

func main() {
	var a [2]int
	a[0]=1
	a[1]=2

	fmt.Println("A: ", a)	

	var b []int
	b = append(b, 3)	

	fmt.Println("B: ", b)	
	
	c := []int{11, 12}
	c = append(c, 13, 14)
	
	fmt.Println("C: ", c)
}
