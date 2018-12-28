package main

import "fmt"

func main() {
	switch {

	case false:
		fmt.Println("false")
	case true:
		fmt.Println("True")
	default:
		fmt.Println("Default")

	}

	x := 5

	switch x {

	case 1:
		fmt.Println("1!")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3!")
	case 5:
		fmt.Println("5!")

	}

}
