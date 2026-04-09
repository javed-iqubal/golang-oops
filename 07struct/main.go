package main

import "fmt"

func main() {

	user := &struct {
		name string
		age  int
	}{
		name: "Maryam",
		age:  7,
	}

	fmt.Printf("Name: %s \nAge: %d \n", user.name, user.age)
}
