// from : https://tutorialedge.net/golang/creating-restful-api-with-golang/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func returnAllArticles(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	//fmt.Fprintf(w, "key : " + key)

	//json.NewEncoder(w).Encode(Articles)

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}

	//if !key {
	//	json.NewEncoder(w).Encode(Articles)
	//}
}

func createNewArticle(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(req.Body)

	var newArticle Article

	json.Unmarshal(reqBody, &newArticle)

	for index, article := range Articles {
		if article.ID == id {
			Articles[index] = newArticle
		}
	}

	json.NewEncoder(w).Encode(newArticle)
}

type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{
			ID : "1",
			Title : "Hello",
			Desc : "Article Description",
			Content : "Article Contents",
		},
		{
			ID : "2",
			Title : "Hello 2",
			Desc : "Article Description",
			Content : "Article Content",
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/articles", returnAllArticles).Methods("GET")
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}