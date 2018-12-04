package main

import (
	"fmt"
)

func main() {
	// result := make(map[string]string)

	// result["status"] = "success"
	// result["message"] = "hello success"
	// result["Status"] = "test"

	result := map[string]string{"status": "success", "message": "hello success", "Status": "test"}

	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Printf("%T\n", len(result))
	fmt.Println(result["status"])
	fmt.Println(result["Status"])

	i := 1
	delete(result, "Status")
	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Printf("%T\n", i)
}
