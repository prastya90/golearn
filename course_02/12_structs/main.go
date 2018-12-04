package main

import (
	"fmt"
	"strconv"
)

type Address struct {
	city, district string
}

type Person struct {
	fullname string
	username string
	age      int
	address  Address
}

func (p Person) greet() string {
	return "Hello my name is " + p.fullname + ". I am live at " + p.address.district + ". I am " + strconv.Itoa(p.age) + " years old."
}

func main() {
	address1 := Address{city: "Jakarta Selatan", district: "Pancoran"}
	person1 := Person{fullname: "Dwi Agung Prastya", username: "helloprastya", age: 25, address: address1}
	// person1 := Person{"Dwi Agung Prastya", "helloprastya", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.fullname)
	// fmt.Println(person1.address.city)
	fmt.Println(person1.greet())
}
