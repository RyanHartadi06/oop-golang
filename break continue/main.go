package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
	fmt.Println("End of loop")

	//break

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}
	fmt.Println("End of loop break")
}
