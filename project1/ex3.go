package main

import "fmt"

type darin int

var x darin
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 43
	fmt.Println(x)
	
	y=int(x)

	fmt.Println(y)
}
