package main

import "fmt"

func main() {
	//	x := type{values} composite literal
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(x)
	z := len(x)

	for i := 0; i<z; i++ {
		x[i]=i+1
	}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i,v)
	}
}
