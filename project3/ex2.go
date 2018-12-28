package main

import "fmt"

const (
	curr = 2018
)

func main() {
	bd := 1998

	for bd <= curr {
		fmt.Println(bd)
		bd++
	}

}
