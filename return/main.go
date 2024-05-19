package main

import "fmt"

func number(a int, b int) int {
	return a + b
}

func addNumber(a int, b int) (result int) {
	result = a + b
	return
}

func getCombine(i string, j int) (number int, text string) {
	number = j + 1
	text = i + " " + "is a string"
	return
}

func main() {
	fmt.Println(number(1, 2))
	fmt.Println(addNumber(1, 2))

	a, b := getCombine("This", 1)
	fmt.Println(a, b)
}
