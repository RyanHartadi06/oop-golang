package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (user User) GetName() string {
	return user.Name
}

func (user User) GetAge() int {
	return user.Age
}

func main() {
	var user User
	user.Name = "John Doe"
	user.Email = "JohnDoe@doe.com"
	user.Age = 25

	name := user.GetName()
	fmt.Println(name)

	age := user.GetAge()
	fmt.Println(age)
}
