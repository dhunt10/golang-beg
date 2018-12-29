package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
	
	for _, v := range arr {
		fmt.Println(v)

	}
}
