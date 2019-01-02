package main

import "fmt"

const (

size =50

)

func main() {
	x := []int{}
	//var size int = 50


	for i := 1; i<=size; i++ {
		x = append(x, i)
		fmt.Println(x)

	}
	
	y := []int{}
	
  	for i := size  ; i>= 1; i-- {

		 
		y = append(x[:i])
		fmt.Println(y)
		fmt.Println(x[:i])	
	}

}
