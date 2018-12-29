package main

import "fmt"

func main() {
	m := map[string][]string{
		"darin": []string {"twenty"}, 
		"cass": []string {"twnetythree"}, 
		"dustin": []string {"twentyfour"}, 
		"donny": []string { "twentysix"},
	}

	fmt.Println(m)

	m["kristen"] = []string{"fiftythree"}	

	fmt.Println(m)
	
	delete(m,"dustin")
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Println(v2)
		}	
	}
}
