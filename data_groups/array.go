package main

import "fmt"

func main() {
	var x [5]int //make sure to specify size
	
	fmt.Println(x)
	
	x[3]= 7
	fmt.Println(x)
	fmt.Println(len(x))
}
