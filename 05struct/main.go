package main

import "fmt"

// Return struct

// annonymous struct return
func getUser() struct {
	name string
	age  int
} {
	return struct {
		name string
		age  int
	}{
		name: "Maria Fatima",
		age:  1,
	}

}

func main() {

	// annonymous struct
	/**

	func functionName() returnType{
		return value
	}

	func functionName() struct{name, age}{

		// function body

		return struct{
			variable
		}{
			initialization
		}
	}

	*/

	person := struct {
		Name   string
		Age    int
		Status bool
	}{
		Name:   "Liza Fatima",
		Age:    3,
		Status: true,
	}

	fmt.Println("Details: ", person)

	fmt.Println(getUser())
}
