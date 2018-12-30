package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

func main() {
	
	p1 := person {
		first:"darin",
		last:"hunt",
		favFlavors: []string{
			"chocolate",
			"vanilla",


		},

	}

	p2 := person {
		first:"cass",
		last:"grace",
		favFlavors: []string{
			"strawberry",
			"coffee",			
		},
	}
	
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
		
	fmt.Println(m)
	fmt.Println(p1)

	
}
