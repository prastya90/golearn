package main

import (
	"fmt"
)

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// // declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitArr := []string{"Apple", "Orange", "Melon", "Kiwi"}

	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:3])
}
