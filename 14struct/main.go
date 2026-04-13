package main

// nil and error example

import "fmt"

func main() {

	var p *int
	fmt.Println(p)

	var s []int
	fmt.Println(s)

	var m map[string]int
	fmt.Println(m)

	var err error
	fmt.Println(err)
}
