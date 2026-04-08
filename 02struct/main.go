package main

import "fmt"

type Product struct {
	Id       string
	Name     string
	Price    float32
	Quantity int
}

// methods of structure
// product Product -> is called receiver - value receiver
func (product Product) calculate() float32 {
	// type conversion (int to float)
	return product.Price * float32(product.Quantity)
}

func main() {

	product := Product{
		Id:       "123",
		Name:     "iPhone",
		Price:    49.5,
		Quantity: 20,
	}

	fmt.Println(product)
	fmt.Println(product.calculate())

}
