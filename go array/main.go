package main

import "fmt"

func main() {
	var a1 = [3]int{1, 2, 3}
	a2 := [...]int{}
	fmt.Println(a2, a1)
}
