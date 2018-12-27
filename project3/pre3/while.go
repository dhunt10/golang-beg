package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	foo()
}

func foo() {
	x := 1

	for {
		if x > 9 {
			break
		}else if x==6 {
			x++
			continue
		}

		fmt.Println(x)
		x++

	}

}
