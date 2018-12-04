package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(greeting("agung"))
	fmt.Println(getSum(2, 5))
}
