package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1><a href=\"/about\">about</a>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1><a href=\"/\">home</a>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting on port 14123")
	http.ListenAndServe(":14123", nil)
}
