package main

import (
	"fmt"
	"struct/composition"
)

func main() {

	// composition.PrintDetails()
	// composition.PrintDetailsUsingPointerToStruct()

	person := composition.CreatePerson(
		"Maryam",
		"Fatma",
		composition.Address{
			City:  "Noida",
			State: "UP",
		},
	)
	fmt.Printf("Name: %s %s \n", person.FirstName, person.LastName)
	fmt.Printf("Address: %s %s \n\n", person.City, person.State)

	person2 := composition.CreatePersonRetunPointer(
		"Maryam",
		"Fatma",
		composition.Address{
			City:  "Noida",
			State: "UP",
		},
	)
	fmt.Printf("Name: %s %s \n", person2.FirstName, person2.LastName)
	fmt.Printf("Address: %s %s \n\n", person2.City, person2.State)

	person3 := composition.PersonFactory()

	fmt.Printf("Name: %s %s \n", person3.FirstName, person3.LastName)
	fmt.Printf("Address: %s %s \n\n", person3.City, person3.State)

}
