package main

import (
	"fmt"
)

func main() {
	ids := []int{1, 2, 3, 3}

	for i, id := range ids {
		fmt.Printf("%d: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Println("ID", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("SUM", sum)

	result := map[string]string{"status": "success", "message": "hello success", "Status": "test"}

	for k, v := range result {
		fmt.Printf("%s: %s\n", k, v)
	}
}
