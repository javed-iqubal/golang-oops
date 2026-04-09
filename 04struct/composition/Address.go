package composition

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	FirstName string
	LastName  string
	Address
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
	fmt.Printf("Name: %s %s \nAddress: %s %s \n\n", person.FirstName, person.LastName, person.Address.City, person.Address.State)
}

func PrintDetailsUsingPointerToStruct() {

	person := &Person{

		FirstName: "Liza",
		LastName:  "Fatma",
		Address: Address{
			City:  "Hyderabad",
			State: "Telangana",
		},
	}
	fmt.Println(&person)
	fmt.Println(person)
	fmt.Printf("Name: %s %s \nAddress: %s %s \n\n", person.FirstName, person.LastName, person.Address.City, person.Address.State)
	//
	fmt.Println(*person)
	fmt.Printf("Name: %s %s \nAddress: %s %s \n\n", (*person).FirstName, (*person).LastName, (*person).Address.City, (*person).Address.State)

}

func CreatePerson(FirstName string, LastName string, Address Address) Person {
	person := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
	return person
}

func CreatePersonRetunPointer(FirstName string, LastName string, Address Address) *Person {
	person := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
	return &person
}

/*
*

	Factory function to initialize the structs with parameter
	Constructors: newFunctionName(args) returnType -> Since Go does not support constructor like other languages, constructors are specified via factory functions
*/
func newPerson(FirstName string, LastName string, Address Address) *Person {
	person := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
	return &person
}

func PersonFactory() *Person {

	person := newPerson(
		"Maryam",
		"Fatma",
		Address{
			City:  "Noida",
			State: "UP",
		},
	)
	return person
}
