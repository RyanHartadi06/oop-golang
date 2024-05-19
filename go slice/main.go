package main

import "fmt"

func main() {

	slice := []string{
		"Azizi",
		"Freya",
		"Gracia",
		"Marsha",
	}

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//change slice to array
	iniArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	iniSlice := iniArray[3:5]

	fmt.Println(iniSlice)
	fmt.Printf("length %d\n", len(iniSlice))
	fmt.Printf("cap %d\n", cap(iniSlice))

	//using make

	//make(type, length, capacity)

	iniMake := make([]int, 5, 5)
	fmt.Println(iniMake)
	fmt.Printf("length %d\n", len(iniMake))
	fmt.Printf("cap %d\n", cap(iniMake))
}
