package main

import "fmt"

const (
	curr = 2018
)

func main() {
	bd := 1998

	for {
		fmt.Println(bd)
		bd++

		if curr == bd {
			break
		}
	}

}
