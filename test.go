package main

import "fmt"

func main() {
	fmt.Println("helo")
	foo()
	loops()
}

func foo() {
	fmt.Println("in the code")
}

func loops() {
	for i := 0; i < 100; i++ {

		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
