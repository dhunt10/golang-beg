package main

import "fmt"

func main() {
	foo()
	bar("James")
	s := woo("Money")
	fmt.Println(s)
	x, y := multwoo("Darin", "Hunt")

	fmt.Println(x, y)
}

func foo() {
	fmt.Println("hellow from foo")
}

func bar(s string) { //accepts one arrameterrrr, gives no return
	fmt.Println("hello,", s)

}

func woo(st string) string { //accepts one parameter and returns a string
	return fmt.Sprint("Hello from woo ", st)

}

func multwoo(s string, o string) (string, bool) { //accepts two parameters and returns two 

	var b bool
	a := fmt.Sprint(s, o, "says hello")
	if s == o {
		b = true

	} else {
		b = false

	}
	return a, b
}
