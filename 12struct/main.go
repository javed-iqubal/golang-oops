package main

import "fmt"

// empty interface - interface without method

var empty interface {
}

func main() {

	empty = 18
	println(empty)
	fmt.Println(empty)

	value, ok := empty.(int)
	if ok {
		println(value)
	}

	empty = "Hello world"
	println(empty)
	fmt.Println(empty)
	str, ok := empty.(string)
	if ok {
		println(str)
	}

	empty = 42.56
	println(empty)
	fmt.Println(empty)
	floatVal, ok := empty.(float64)
	if ok {
		println(floatVal)
	}

}
