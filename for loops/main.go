package main

func main() {
	// var items = []string{"a"}
	// var items2 = []string{"f", "g"}

	// for i := 0; i < len(items); i++ {
	// 	for j := 0; j < len(items2); j++ {
	// 		println(items[i] + items2[j])
	// 	}
	// }

	//using range

	cars := []string{"Volvo", "Toyota", "Ford"}

	for index, car := range cars {
		println(index, car)
	}

	person := map[string]string{
		"name": "John Doe",
		"job":  "Software Engineer",
	}

	for _, value := range person {
		println(value)
	}
}
