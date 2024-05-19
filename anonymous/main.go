package main

import "fmt"

func main() {

	getHello := func() {
		fmt.Println("Hello")
	}

	getHello()

	//with params
	getPeople := func(name string, age int) {
		fmt.Print("Name: ", name, " Age: ", age, "\n")
	}

	getPeople("John Doe", 27)

	//with return
	getTotal := func(a int, b int) int {
		return a + b
	}

	total := getTotal(1, 2)
	fmt.Println(total)
}
