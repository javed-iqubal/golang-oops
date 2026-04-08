package main

import "fmt"

// class declaration structure

type Person struct {
	Name  string
	Age   int
	ID    string
	Email string
}

func main() {

	// Basic instance creation
	var person Person
	fmt.Println(person)
	fmt.Println()

	person.Name = "Javed Iqubal"
	person.Age = 40
	person.ID = "AUG07"
	person.Email = "test@gmail.com"

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Println("ID: ", person.ID)
	fmt.Println("Email: ", person.Email)

	fmt.Println(person)

	// Initialize using literal pattern
	var person2 = Person{
		Name:  "Liza",
		Age:   40,
		ID:    "SEP10",
		Email: "liza@gmail.com",
	}

	fmt.Println(person2)

	// without field names
	var person3 = Person{
		"Maria",
		1,
		"JAN13",
		"maria@gmail.com"}
	fmt.Println(person3)
}
