package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // creates an array w 2 empty slots and capacity of 5
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Greg")

	fmt.Println(userNames)

	// courseRatings := map[string]float64{}
	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	// fmt.Println(courseRatings)
	courseRatings.output()
}
