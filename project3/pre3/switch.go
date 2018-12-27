package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("False")
	case (2 == 4):
		fmt.Println("This should not print")
	case (3 == 3):
		fmt.Println("That is correct")
		fallthrough //more than one thing can be corret , doestnt automatically break
	case (4==4):
		fmt.Println("also true")
	default: 
		fmt.Println("nothing was true")
	}
//generally dont use fallthrough
}
