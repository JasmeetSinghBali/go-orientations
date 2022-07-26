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
	"time"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       int
	Data      BookTicket
	TimeStamp string
	Hash      string
	PrevHash  string
}

// 📝 used by newBook /new POST
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

// 📝 used by writeBlock / POST
type BookTicket struct {
	BookID     string `json:"book_id"`
	User       string `json:"user"`
	BoughtDate string `json:"bought_date"`
	IsGenesis  bool   `json:"is_genesis"`
}

// 🎈 This can be replaced by a DB, for simplicity purpose it we are considering a slice of type Block
type Blockchain struct {
	// slice of type of Block
	blocks []*Block
}

var blockchain *Blockchain

// @desc - struct method for Block struct to generate new hash for newly created block
func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)

	// data to create the hash
	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash

	hash := sha256.New()
	hash.Write([]byte(data))

	// attach hash to the new block
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

// @desc - CreateBlock , create a new block with bookTicket and prevBlock data
func CreateBlock(prevBlock *Block, bookTicket BookTicket) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.PrevHash = prevBlock.Hash
	block.Data = bookTicket
	block.TimeStamp = time.Now().String()
	block.generateHash()

	return block
}

// @desc- AddBlock of type Blockchain struct
// syntaxHelp-  ()structMethodName()
func (bc *Blockchain) AddBlock(data BookTicket) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(prevBlock, data)

	// check validity of the block
	// if valid then append new block to blockchain
	if validBlock(newBlock, prevBlock) {
		bc.blocks = append(bc.blocks, newBlock)
	}
}

// @desc- check newly created block is valid with reff to prevHash and Hash
func validBlock(block, prevBlock *Block) bool {
	// check wheather current block holds the hash in PrevHash same as prevBlock.Hash value
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	// validate hash
	if !block.validateHash(block.Hash) {
		return false
	}
	// check position
	if prevBlock.Pos+1 != block.Pos {
		return false
	}

	return true
}

// @desc- struct method validate hash of the newly created block
func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

// @desc - handler for writing block to blockchain / POST
func writeBlock(w http.ResponseWriter, r *http.Request) {
	var bookTicket BookTicket
	// create a BookTicket
	if err := json.NewDecoder(r.Body).Decode(&bookTicket); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("BookTicketCreationError: %v", err)
		w.Write([]byte("could not create bookTicket"))
		return
	}
	// pass the BookTicket now as data in the block to create new block
	// and then add the block to the blockchain
	blockchain.AddBlock(bookTicket)
	resp, err := json.MarshalIndent(bookTicket, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// @desc - handler for creating new book via /new
func newBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	// 📝convert the json new book payload into Book struct
	// decoding or unmarshalling and storing in book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failedError: %v", err)
		w.Write([]byte("could not create a new book"))
		return
	}

	// create & assign book ID to book created
	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	// send the newly created book back to user in response as json after marshaling it
	newBookMarshaled, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("CreationError: %v", err)
		w.Write([]byte("could not save book data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(newBookMarshaled)
	return
}

// create a genesis block
func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookTicket{IsGenesis: true})
}

// create new blockchain with genesisblock
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// @desc- get the entire exisitng blockchain
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	blockchainBytes, err := json.MarshalIndent(blockchain.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	io.WriteString(w, string(blockchainBytes))
}

// @desc - entry_point for program
func main() {

	// instantiating a blockchain
	blockchain = NewBlockchain()
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET") // get entire blockchain
	r.HandleFunc("/", writeBlock).Methods("POST")   // write on the block @returns BookID that is getting purchased and is passed to BookTicket
	r.HandleFunc("/new", newBook).Methods("POST")   // to create new book

	// sub goroutine
	go func() {
		for _, block := range blockchain.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}
