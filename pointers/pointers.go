package main

import "fmt"

func main() {
	age := 32

	var agePointer *int // this is a pointer
	agePointer = &age   // this is getting the pointer to the age variable

	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult Years:", age)

	// adultYears := getAdultYears(agePointer)
	// fmt.Println("Adult Years:", adultYears)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
