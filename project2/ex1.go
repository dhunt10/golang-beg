package main

import "fmt"

var darin int = 5
var space string
func main() {
	//	foo()
	number()
	compare()
}

func foo() {
	for i := 0; i < 1000; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			fmt.Println("Divisible by 7")
		} else {
			darin = darin + 5
		}
		fmt.Printf("not  ")
		fmt.Println(darin)
	}
}

func number() {
	x := 5
	fmt.Printf("%v\t%b\t%#x\n", x, x, x)
}

func compare() {
	a :=(42 == 42)
	b :=(42 <= 43)
	c :=(42 >= 43)
	d :=(42 != 43)
	e :=(42 < 43)
	f :=(42 > 43)

	fmt.Println(a,b,c,d,e,f)
}
