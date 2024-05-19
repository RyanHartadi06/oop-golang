package main

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func anyParams(name string, data string, nums ...int) {
	println(name, data)
	for _, num := range nums {
		println(num)
	}
}

func main() {

	//normal params
	sumAll := sum(1, 2, 3, 4, 5)
	println(sumAll)

	//with any params
	anyParams("John Doe", "Data", 1, 2, 3, 4, 5)

	// with slice
	slice := []int{1, 2, 3, 4, 5}
	total := sum(slice...)
	println(total)

}
