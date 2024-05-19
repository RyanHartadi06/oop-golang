package main

func message(a map[string]int, b int) {
	for key, value := range a {
		if value == b {
			println(key, value)
		}

	}
}

func main() {

	maps := map[string]int{
		"data": 1,
		"info": 2,
	}

	message(maps, 2)
}
