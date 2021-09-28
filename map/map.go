package main

import "fmt"

func main() {
	mymap := make(map[string]string)

	mymap["key1"] = "value1"
	mymap["key2"] = "value2"

	fmt.Println(mymap)
	fmt.Println(mymap["key1"])

	mymap["key3"] = "value3"
	fmt.Println(mymap)

	delete(mymap, "key3")
	fmt.Println(mymap)

	//maps can be declared like this too
	mymap2 := map[string]string{
		"key1": "mapped value1",
		"key2": "mapped value2",
	}

	fmt.Println(mymap2)
	fmt.Println((mymap2["key2"]))
}
