package main

import "fmt"

func main()  {
	defer foo()
	defer bar("bar")

	s1 := returnString("World")
	fmt.Println(s1)

	//s, err := checkFoo("foo")
	//
	//fmt.Println(s, err)
}

func foo() {
	fmt.Println("func foo")
}

func bar(bar string) {
	fmt.Printf("Hello %s\n", bar)
}

func returnString(s string) string {
	return "Hello " + s
}

func checkFoo(s string) (string, bool) {
	a := s
	if a == "foo" {
		return a, true
	}

	return a, false
}