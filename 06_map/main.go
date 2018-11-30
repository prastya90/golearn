package main

import "fmt"

func main()  {
	m := map[string]int {
		"status": 1,
		"msg": 2,
	}

	fmt.Println(m)
	fmt.Println(m["status"])

	v, err := m["status"]
	fmt.Println(v)
	fmt.Println(err)

	m["test"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	m["test"] = 22

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "test")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
