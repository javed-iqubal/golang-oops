package main

import "fmt"

type Customer struct {
	FirstName string
	LastName  string
}

// value receiver: never updates the instance members
func (customer Customer) updateName() {
	customer.FirstName = "Javed"
	customer.LastName = "Iqubal"

	// fmt.Println(customer)
}

// pointer/reference receiver: updates instances members, because pointers share memory
func (customer *Customer) updateNameUsingReference() {
	customer.FirstName = "Javed"
	customer.LastName = "Iqubal"

	fmt.Println(customer, "  ", *customer)
}

func main() {

	customer := Customer{}
	fmt.Println("Before updateName")
	fmt.Println(customer.FirstName, customer.LastName)
	customer.updateName()
	fmt.Println("After updateName")
	fmt.Println(customer.FirstName, customer.LastName)

	customer.updateNameUsingReference()
	fmt.Println("After updateNameUsingReference")
	fmt.Println(customer.FirstName, customer.LastName)

}
