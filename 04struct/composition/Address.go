package composition

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

func PrintDetails() {

	person := Person{

		FirstName: "Maria",
		LastName:  "Fatma",
		Address: Address{
			City:  "New Delhi",
			State: "Delhi",
		},
	}
	fmt.Println(person)
	fmt.Printf("Name: %s %s \nAddress: %s %s \n", person.FirstName, person.LastName, person.Address.City, person.Address.State)
}
