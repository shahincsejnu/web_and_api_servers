package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func main() {
	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}