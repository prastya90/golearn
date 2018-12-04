package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Articles []Article

var articles = Articles{
	Article{Id: "1", Title: "Title 1", Content: "Content 1", Author: "prastya"},
	Article{Id: "2", Title: "Title 2", Content: "Content 2", Author: "prastya"},
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article Article

	for i := range articles {
		if articles[i].Id == key {
			article = articles[i]
			break
		}
	}

	if (Article{} != article) {
		json.NewEncoder(w).Encode(article)
		return
	}

	json.NewEncoder(w).Encode(nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func handleRequest() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/articles", getAllArticles)
	muxRouter.HandleFunc("/articles/{id}", getArticle)

	log.Fatal(http.ListenAndServe(":3000", muxRouter))
}

func main() {
	handleRequest()
}
