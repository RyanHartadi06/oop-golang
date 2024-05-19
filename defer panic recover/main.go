package main

import "fmt"

//defer
func txtNumber() {
	fmt.Println("tiga")
}

//panic

func division(a, b int) {
	defer recoverHandler()
	if b == 0 {
		panic("Pembagi tidak boleh 0")
	} else {
		result := a / b
		fmt.Println(result)
	}
}

//recover
func recoverHandler() {
	err := recover()

	if err != nil {
		fmt.Println("Error dengan pesan", err)
	}
}

func main() {
	defer txtNumber()

	fmt.Println("satu")
	fmt.Println("dua")

	division(10, 0)
}
