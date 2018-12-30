package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	lux bool
} 

func main() {
	
	t := truck{
		vehicle: vehicle{
			doors:2,	
			color: "white",
		},
		fourWheel: true,

	}
	s := sedan{
		vehicle: vehicle{
			doors:4,
			color: "black",
		},
		lux: true,
	}	
	fmt.Println(t,s)
}
