package main

import "fmt"

func main() {
	// var name = "Agung"
	var age = 20
	const isCool = true

	// name := "Test"
	price := 10.2
	name, email := "agung", "agung@gmail.com"

	fmt.Println(name, age, isCool, price, email)
	fmt.Printf("%T\n", price)
}
