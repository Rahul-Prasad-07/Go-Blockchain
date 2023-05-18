package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// we are creating four structs for our blockchain which are Block, BookCheckout, Book and Blockchain
// It is used to store the data of the blockchain and the books

// Block is used to store the data of the block
// for Block we're not going to have any json tags beacuse we are not going to send this data to the user
// we don't have any apis for our project, we are just going to use this json data for sending to postman for testing but block is created by itself when we hit new in route and it will store the data of bookcheckout struct in it.
type Block struct {
	Position  int
	Data      BookCheckout
	TimeStamp string
	Hash      string
	PrevHash  string
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

// why we are not using any databse : Blockchain is itself a database. It's store the transactional information.
// Block is created and store in this blockchain variable
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

// go mod tidy - to install all the dependencies
// go get -u github.com/gorilla/mux - to install gorilla mux
// go run main.go - to run the project
// go build - to build the project
// ./blockchain - to run the build project
// go mod init github.com/yourusername/blockchain - to create a module
