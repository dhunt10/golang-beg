package main

import "fmt"

func main() {
	x := "hello"
	switch x {
	case "bye", "hola", "goodbye":
		fmt.Println("not hello")
	case "hello", "yo":
		fmt.Println("correct")
	default:
		fmt.Println("default used")
	}

}
