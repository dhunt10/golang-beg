package main

import "fmt"

func main() {
	//	x := type{values} composite literal

	x := make([]int, 10, 100) //type, length, capacity
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 2000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//matrix

	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 0}
	xp := [][]int{a, b}

	fmt.Println(xp)
}
