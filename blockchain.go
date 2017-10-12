// Author: Dhriti Shikhar
// Implementing blockchain in Golang
// Date: October 3, 2017
// Email: dhriti.shikhar.rokz@gmail.com

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type blockchain struct {
	chain               []block       // store blockchain
	currentTransactions []Transaction // store transactions
	lastBlock           block
}

type Transaction struct {
	Sender   string `json:"sender"`
	Reciever string
	Amount   int
}

type block struct {
	index        int
	timestamp    time.Time
	transactions []Transaction
	proof        int
	previousHash string // this creates a chain and gives blockchain immutability
}

type blockchainer interface {
	mine() block
	newTransaction(newTransaction Transaction) int
	fullChain()
	registerNode()
	consensus()
	proofOfWork()
	newBlock(proof int, previousHash string) block
	hash()
	lastBlock()
}

// Validates the Proof Of Work
func validProof(prevProof int) bool {
	pp := []byte(strconv.Itoa(prevProof))
	gHash := sha256.Sum256(pp)
	guessHash := string(gHash[:])
	return guessHash[:4] == "0000"
}

// Let p be proof of previous block
// Let p' be the current proof which we are trying to figure out
//
// Proof of Work Rule
// ------------------
// Find p' such that hash(p*p') contains leading 4 zeroes
func (b blockchain) proofOfWork(prevProof int) int {
	currentProof := 0
	for !validProof(prevProof) {
		currentProof += 1
	}
	return currentProof
}

func mine(w http.ResponseWriter, r *http.Request) {
	b := blockchain{}
	b.mine()
}

func (b blockchain) mine() block {
	// get the last block's proof
	lastBlock := b.lastBlock
	lastProof := lastBlock.proof

	// get the proof for current block
	currentProof := b.proofOfWork(lastProof)

	// node identifier
	nodeID := uuid.NewV4().String()
	newTransac := Transaction{
		Sender:   "0", // sender is "0" to signify new transaction
		Reciever: nodeID,
		Amount:   1,
	}
	b.newTransaction(newTransac)

	// add a new block to the chain
	block := b.newBlock(currentProof, "0")

	return block
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	b := blockchain{}

	var t Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err)
	}

	//TODO: check if all the required fields are present in JSON body

	output := []byte(strconv.Itoa(b.newTransaction(t)))
	w.Write(output)
}

// Creates a new transaction
// Adds it to the list of transactions under block
func (b blockchain) newTransaction(newTransaction Transaction) int {
	var t []Transaction
	t = append(t, newTransaction)

	b = blockchain{
		currentTransactions: t,
	}
	return b.getLastBlockIndex() + 1 // returns the index of the next block to be mined
}

func (b blockchain) getLastBlockIndex() int {
	return len(b.chain) - 1
}

func fullChain(w http.ResponseWriter, r *http.Request) {
	b := blockchain{}
	output := []byte(strconv.Itoa(len(b.chain)))
	w.Write(output)
}

func registerNode(w http.ResponseWriter, r *http.Request) {
	/*b := blockchain{}

	// get nodes
	var t Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err)
	}

	output := []byte(strconv.Itoa(len(b.chain)))
	w.Write(output)*/
}

// Creates a new block i.e. genesis block in the Blockchain
// This is called when blockchain is instantiated
func (b blockchain) newBlock(proof int, previousHash string) block {

	newBlockData := block{
		index:        len(b.chain) + 1,
		timestamp:    time.Now(),
		transactions: b.currentTransactions,
		proof:        proof,
		previousHash: previousHash,
	}

	// reset current list of transactions
	b.currentTransactions = b.currentTransactions[:0]
	b.chain = append(b.chain, newBlockData)
	return newBlockData
}

// hashes a block
func hash(block) {
}

// returns the last block in the chain
func lastBlock() {
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintf("Welcome")
	})
	http.HandleFunc("/mine", mine)

	http.HandleFunc("/transactions/new", newTransaction)

	http.HandleFunc("/chain", fullChain)

	http.HandleFunc("/nodes/register", registerNode)

	//http.HandleFunc("/nodes/resolve", consensus)
	http.ListenAndServe(":9090", nil)
}
