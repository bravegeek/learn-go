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
	hobbies4 := hobbies[1:]
	fmt.Println(hobbies4)

	fmt.Println(5)
	goals := []string{"learn-go", "finish-course"}
	fmt.Println(goals)

	fmt.Println(6)
	goals[1] = "complete-it"
	goals = append(goals, "be-awesome-at-go")
	fmt.Println(goals)

	fmt.Println(7)
	products := []Product{{title: "N5105", id: "1", price: 125.32}, {title: "Jonsbo N1", id: "2", price: 139.99}}
	products = append(products, Product{title: "SFX Power Supply", id: "32", price: 99.99})

	fmt.Println(products)
}
