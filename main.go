package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "encoding/hex"
	"time"
	// "crypto/sha256"

	"github.com/gorilla/mux"
)

// we are creating four structs for our blockchain which are Block, BookCheckout, Book and Blockchain
// It is used to store the data of the blockchain and the books

// Block is used to store the data of the block
// for Block we're not going to have any json tags beacuse we are not going to send this data to the user
// we don't have any apis for our project, we are just going to use this json data for sending to postman for testing but block is created by itself when we hit new in route and it will store the data of bookcheckout struct in it.
type Block struct {
	Pos       int
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
	IsGenesis        bool   `json: "is_genesis"`
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

// --> Here we are passing Book-id, Checkout Data
func writeBlock(w http.ResponseWriter, r *http.Request) {

	var checkoutitem *BookCheckout

	// decode the jason data and store it in checkoutitem
	if err := json.NewDecoder(r.Body).Decode(&checkoutitem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not decode the request body to json: %v", err)
		w.Write([]byte("could not create a new block"))
		return
	}

	// if there is no error then we are going to create a new block
	Blockchain.AddBlock(checkoutitem) // create new func AddBlock, which will add the block in the blockchain

}

func (bc *Blockchain) AddBlock(data BookCheckout) {

	// we are going to create a new block and store the data of book in it
	// we need prevHash for the new block, so we are going to get the prevHash from the last block
	// we are going to create a new block by passing prevBlock and data
	// we are going to append the new block in the blockchain by validating the new block

	prevBlock := bc.blocks[len(bc.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}

	// now we are going to createBlock func for creating a new block
	// now we are going to validBlock func for validating the new block
}

func validBlock(block, prevBlock *Block) bool {

	// first we check the hash of the prevBlock and the prevHash of the new block, if it's not same then we return false
	// then we validate the hash of block by calling validateHash func
	// then we check the position of the new block, if it's not same as the prevBlock then we return false

	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}

	// if all the above condition is true then we return true
	return true

}

// create a new func validateHash(struct method)
func (b *Block) validateHash(hash string) bool {

	// here we call generateHash func to generate the hash of the new block
	// then we check the hash of the new block and the hash of the new block, if it's not same then we return false

	b.generateHash()
	if b.Hash != hash {
		return false
	}

	return true

}

func CreateBlock(prevBlock *Block, data BookCheckout) *Block {

	// first we define block variable with empty Block struct
	// then we are going to create a new block by passing prevBlock and data and other properties of block

	block := &Block{}

	block.Pos = prevBlock.Pos + 1
	block.TimeStamp = time.Now().String()
	block.PrevHash = prevBlock.PrevHash
	block.generateHash() // we have to create generateHash func and generate has for this block

	return block

}

func (b *Block) generateHash() {

	// we have data in block, we have to convert this data into hash
	// first we have to convert the Data into json and store it into bytes :--> this is alredy exsisiting data in block
	// then we have to create a new hash by storing the data & properties of block in one varibale data by adding with string
	// now we are using sha256 to create a new hash and store it into b

	bytes, _ := json.Marshal(b.Data)

	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))

}

// --> Here you get book json data with id & It will create new block and store the data of book in it
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

// in Func main ()
// --> as soon as you run the project, it will create a new blockchain : Blockchain = NewBlockchain()
// now you have to create NewBlockchain func, it dosent take any argument and it returns &blockchain(which has multiple blocks) and
//  we have slice of blocks but the first block is genesis block(which is the first block in the blockchain) and now create that func GenesisBlock()

func NewBlockchain() *Blockchain {

	return &Blockchain{[]*Block{GenesisBlock()}}
}

// --> now we have to create GenesisBlock func, it dosent take any argument and it returns genesis &block
func GenesisBlock() *Block {

	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func main() {

	Blockchain = NewBlockchain()

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
