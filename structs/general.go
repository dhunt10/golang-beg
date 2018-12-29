package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first_name: "don",
			last_name:  "hunt",
			age:        63,
		},
		ltk: true,
	}

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

	p3 := person{
		first_name: "dustin",
		last_name:  "hunt",
		age:        24,
	}
	fmt.Println(p1, p2, p3, sa1)
	fmt.Println(p1.first_name, p1.age)
}
