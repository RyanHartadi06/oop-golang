package main

import "fmt"

func recursive(x int) int {
	if x == 11 {
		return 0
	}

	fmt.Println(x)

	return recursive(x + 1)
}

func main() {
	recursive(1)
}
