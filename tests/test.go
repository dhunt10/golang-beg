package main

import "fmt"

var y = 17 //can declare a variable outside of a function body, it is global

var z int //z is type int but not yet defined, it will be default to 0

func main() {
	x := 42 //assigns and declares
	x = 99  //assigns a varaible that has already been delcared
	fmt.Println(x)
	fmt.Println("helo")
	foo()
	//	loops()
	strings()
	makeTypes()
}

func foo() {
	fmt.Println("in the code")
}

func loops() {
	for i := 0; i < 100; i++ {

		if i%3 == 0 {
			fmt.Println("divisible by three", i)
			fmt.Printf("%T\n", i) //prints type you can also replace T for b fro binary, x for hex, #x for o x hex
			//     \n is new line    \t is tab for printing      \v is vertical tab
		} else if i%2 == 0 {

			fmt.Println("divisible bt two", y, " ", i)
		} else if i%20 == 0 {
			fmt.Println("divisible by 20", z)
		}

	}
}

func strings() {

	r := "darin hunt" //string

	fmt.Println(r)
	fmt.Printf("%b\t%#x\t%o\n", y, y, y)
	Value := fmt.Sprintf("%b\t%#x\t%o", y, y, y) //assigbns value, a string to the values
	fmt.Println(Value)                           //prints the value
}

func makeTypes() {
	type hotdog int
	var b hotdog = 43
	var a int
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

