package main

import "fmt"

func main() {
	x := 42

	if x == 40 {
		fmt.Println("value is 40")
	} else if x == 41 {
		fmt.Println("value is 41")
	} else if x == 42{
		fmt.Println("value is 42")
	} else {
		fmt.Println("Value is not 40,41 or 42")
	}

}
