package main

import "fmt"

const (
	a = 1
	b = 2.3454
	c = "hello"

	d = iota
	e = iota
	f = iota
)

const (
	g = iota
	h = iota
	i = iota



)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
