package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		rem := i % 4
		fmt.Println(rem)
		if rem==0 {
			fmt.Println("Divisible by 4")
		
		}else if rem==1 {
			fmt.Println("Remainder is 1")
		}else {
			fmt.Println("Other activated")
	
		}
	}

}
