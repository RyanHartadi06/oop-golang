package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice[3])
	slice[2] = 10
	fmt.Println(slice)
}
