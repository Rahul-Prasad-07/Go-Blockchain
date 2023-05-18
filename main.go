package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// we are creating four structs for our blockchain which are Block, BookCheckout, Book and Blockchain
// It is used to store the data of the blockchain and the books

// Block is used to store the data of the block
type Block struct {
}

// BookCheckout is used to store the data of the book checkout
type BookCheckout struct {
	BookID           string `json:"book_ID"`
	User             string `json:"user"`
	BookcheckoutDate string `json:"bookcheckout_date"`
	IsGenesis        string `json: "is_genesis"`
}

// Book is used to store the data of the book
// when we hit new in route, It will create new book. but all the details of book will be store in the bookcheckout struct.
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publishDate"`
	ISBN        string `json:"isbn:"`
}

// Blockchain is used to store the data of the blockchain
type Blockchain struct {
	blocks []*Block
}

var blockchain *Blockchain //global variable, we are going to use this to store our blockchain

// we are creating three functions for our blockchain which are getBlockchain, writeBlock and newBook

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
