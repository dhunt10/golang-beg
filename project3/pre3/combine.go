package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}
		if i%11 == 0 {
			fmt.Printf("Value %d is divisible by ten\n", i)
		} else {
			continue
		}

	}

}
