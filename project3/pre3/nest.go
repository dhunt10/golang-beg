package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("The inner loop : %d\t The Outer loop: %d\n ", i, j)
		}

	}

}
