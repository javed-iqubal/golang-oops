package main

import "fmt"

type Employee struct {
	Name string

	// Nested structure
	Address struct {
		City  string
		State string
	}
}

func main() {

	var emp Employee

	emp.Name = "Mike Witte"
	emp.Address.City = "Pune"
	emp.Address.State = "Maharastra"

	fmt.Println(emp)

	fmt.Println(emp.Name, emp.Address.City, emp.Address.State)
}
