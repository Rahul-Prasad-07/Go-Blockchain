package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "encoding/hex"
	// "time"
	// "crypto/sha256"

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

func newBook(w http.ResponseWriter, r *http.Request) {

	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create new book: %v", err)
		w.Write([]byte("could not create new book"))
		return
	}

	//if there is no error then we are going to create a new book

	// what this md5.new() is doing is it's creating a new hash
	// what this line is doing is it's writing the string of book.ISBN+book.PublishDate to the hash
	// what this line is doing is it's converting the hash to a string and storing it in the book.ID by using fmt.Sprintf("%x", h.Sum(nil))
	// what this md5.new() is doing is it's creating a new hash

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	// you have created the book and now you want to share that with user so you have to send the response back to the user
	// so we are going to create a json response and send it back to the user

	// what this line means is it's going to take the book struct and convert it into a json response. and it's going to indent it with "" and it's going to use space as a separator
	// if there is an error then we are going to send the error back to the user

	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal book payload to json: %v", err)
		w.Write([]byte("could not save the book data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	// What's happening here : you have to send that created book as responseto the user. so for that u have convert in into the json response and send it back to the user
	// after creating book, we are going to create a block and store the data of book in it: so the next func is writeBlock

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
