package main

import "fmt"

func main() {
	a := 5
	b := &a

	*b = 2
	c := 3
	fmt.Println(a, &a, b, *b, c, &c)
	b = &c
	*b = 10

	fmt.Println(a, &a, b, *b, c, &c)

}
