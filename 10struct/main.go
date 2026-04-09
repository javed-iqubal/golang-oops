package main

import "fmt"

// interface declaration
/*
	type InrefaceName interface{

		// method declartion
	}
*/

// interface declaration
type Speaker interface {
	// method declaration
	Speak()
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Dog Barks")
}

type Cat struct {
}

func (c Cat) Speak() {

	fmt.Println("Cts mew")
}

func main() {

	var speaker Speaker
	speaker = Dog{}
	speaker.Speak()

	speaker = Cat{}
	speaker.Speak()
}
