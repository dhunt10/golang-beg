package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	//	fmt.Println(m[32])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	//looping through a map

	m["todd"] = 33//adds to the map

	fmt.Println(m)
	
	delete(m,"James")
	
	if v, ok := m["Greg"]; ok {

		fmt.Println("itys in there", v)
	}else{
		fmt.Println("Not in there")

	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
