package main

import "fmt"

func main() {
	//	x := type{values} composite literal
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(x)

	x = append(x,77,88,99,104)
	fmt.Println(x)
	
	y := []int{1,2,3,4,5}
	x = append(x,y...) //putting y into x, '...' allows for all values to be sent
	fmt.Println(y)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
