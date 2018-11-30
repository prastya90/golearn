package main

import "fmt"

var d string = "global"

type bar int

var j bar = 1

var k = j


func main()  {
	var a int = 2
	var b = "test"
	c := 0.4


	fmt.Printf("%T: ", a)
	fmt.Println(a)

	fmt.Printf("%T: ", b)
	fmt.Println(b)

	fmt.Printf("%T: ", c)
	fmt.Println(c)

	c = 1 // why c can set integer format

	fmt.Printf("%T: ", a)
	fmt.Println(a)

	fmt.Printf("%T: ", c)
	fmt.Println(c)

	fmt.Println(d)

	foo()

	fmt.Println(d)

	d += " add test"

	fmt.Println(d)

	var e int32

	fmt.Println(e)

	var f float32

	fmt.Println(f)

	var g string

	fmt.Printf("g = %s => it's empty string %t\n", g, g == "")

	var h int = 9
	fmt.Printf("%d\n", h)
	fmt.Printf("%b\n", h)

	fmt.Printf("%T %d\n", j,j)
	fmt.Printf("%T %d\n", k,k)

	c = 1.1
	a = int(c)
	fmt.Printf("%T %d\n", a,a)

	l := fmt.Sprintf("%f", c)
	fmt.Printf("%T %s\n", l,l)
}

func foo()  {
	d = "global change"
}
