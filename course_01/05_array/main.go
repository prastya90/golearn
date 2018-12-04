package main

import "fmt"

func main() {
	x := []int{}

	fmt.Println(x)

	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y)

	fmt.Println(len(y))

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println(y[1:])
	fmt.Println(y[1:3])

	x = append(x, 20, 22, 23)

	fmt.Println(x)

	y = append(y[0:], x[1:]...)

	fmt.Println(y)

	z := make([]int, 10, 20)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}
