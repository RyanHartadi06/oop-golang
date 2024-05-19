package main

import "fmt"

func main() {
	mapping := map[string]string{
		"name": "John Doe",
		"age":  "27",
	}

	mapping_make := make(map[string]string)
	mapping_make["name"] = "John Doe"
	mapping_make["age"] = "27"

	fmt.Println(mapping)
	fmt.Println(mapping["name"])

	fmt.Println(mapping_make)
	fmt.Println(mapping_make["name"])

	//empty map
	var mapping_empty map[string]string
	var mapping_make_empty = make(map[string]string)
	fmt.Println(mapping_empty)
	fmt.Println(mapping_make_empty)
}
