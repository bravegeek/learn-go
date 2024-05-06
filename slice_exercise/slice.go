package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"running", "backpacking", "scuba-diving"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])

	fmt.Println(2)
	expensiveHobbies := hobbies[1:]
	fmt.Println(expensiveHobbies)

	fmt.Println(3)
	// hobbies3 := hobbies[:2]
	hobbies3 := hobbies[0:2]
	fmt.Println(hobbies3)

	fmt.Println(4)
	// hobbies3 = hobbies3[1:] // this will only return "backpacking" because of the existing slice
	hobbies3 = hobbies3[1:3] // returns "backpacking" and "scuba-diving", uses the backing array to get the rest
	fmt.Println(hobbies3)

	fmt.Println(5)
	goals := []string{"learn-go", "finish-course"}
	fmt.Println(goals)

	fmt.Println(6)
	goals[1] = "complete-it"
	goals = append(goals, "be-awesome-at-go")
	fmt.Println(goals)

	fmt.Println(7)
	products := []Product{
		{title: "N5105", id: "1", price: 125.32},
		{title: "Jonsbo N1", id: "2", price: 139.99},
	}
	products = append(products, Product{"SFX Power Supply", "32", 99.99})

	fmt.Println(products)
}
