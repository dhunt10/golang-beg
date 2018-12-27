package main

import "fmt"

var darin int = 5

func main() {
	fmt.Println("Hello")
	fmt.Println(darin)
	foo()
}

func foo() {
	for i := 0; i < 1000; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			fmt.Println("Divisible by 7")
		} else {
			darin = darin + 5
		}
		fmt.Printf("not  ")
		fmt.Println(darin)
	}
}
