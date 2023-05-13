package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
}

type BookCheckout struct {
}

type Book struct {
}

type Blockchain struct {
	blocks []*Block
}

var blockchain *Blockchain //global variable, we are going to use this to store our blockchain

func getBlockchain() {

}

func writeBlock() {

}

func newBook() {

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	log.Println("listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))

}
