package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	for _, v := range arr {
		fmt.Println(v)

	}
}
