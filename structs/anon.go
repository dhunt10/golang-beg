package main

import "fmt"

func main() {
	p1:= struct{
		first string
		last string
		age int
	}{
		first:"darin",
		last:"hunt",
		age:20,
	}
	fmt.Println(p1)
	
}
