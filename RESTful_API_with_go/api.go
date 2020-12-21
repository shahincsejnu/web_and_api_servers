// from : https://tutorialedge.net/golang/creating-restful-api-with-golang/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func returnAllArticles(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{
			Title : "Hello",
			Desc : "Article Description",
			Content : "Article Contents",
		},
		{
			Title : "Hello 2",
			Desc : "Article Description",
			Content : "Article Content",
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage)
	router.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":8080", router))
}