package main

import "fmt"

func main() {
	slc1 := []string{"darin","hunt","is","twenty"}
	slc2 := []string{"cassandra","grace","is","twenty three"}
	fmt.Println(slc1)
	fmt.Println(slc2)
	
	slc3 := [][]string{slc1,slc2}

	fmt.Println(slc3)

}
