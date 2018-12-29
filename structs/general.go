package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {

	p1 := person{
		first_name: "darin",
		last_name:  "hunt",
		age:        20,
	}

	p2 := person{
		first_name: "cass",
		last_name:  "hunt",
		age:        23,
	}

	fmt.Println(p1, p2)

}
