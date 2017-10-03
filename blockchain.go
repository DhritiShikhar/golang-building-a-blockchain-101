// Author: Dhriti Shikhar
// Implementing blockchain in Golang
// Date: October 3, 2017
// Email: dhriti.shikhar.rokz@gmail.com

package main

import (
	"crypto/sha256"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type blockchain struct {
	chain               []string // store blockchain
	currentTransactions []string // store transactions
	lastBlock           block
}

type transaction struct {
	sender    string
	recipient string
	amount    int
}

type block struct {
	index        int
	timestamp    time.Date
	transaction  transation
	proof        int
	previousHash string // this creates a chain and gives blockchain immutability
}

type blockchainer interface {
	mine()
	newTransaction(newTransaction transaction) int
	fullChain()
	registerNode()
	consensus()
	proofOfWork()
}

// Validates the Proof Of Work
func validProof(prevProof) bool {
	guessHash := sha256.Sum256(prevProof)
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
	for !validProof(lastProof) {
		currentProof += 1
	}
	return currentProof
}

func (b blockchain) mine() {
	// node identifier
	nodeID := uuid.NewV4()
	u2, err := uuid.FromString(nodeID)

	// get the last block's proof
	lastBlock := b.lastBlock
	lastProof := lastBlock.Proof

	// get the proof for current block
	currentProof := b.proofOfWork(lastProof)

	newTransac := newTransaction{
		sender:   "0",
		reciever: nodeID,
		amount:   1,
	}
}

// Creates a new transaction
func (b blockchain) newTransaction(newTransaction transaction) int {
	b := append(b, newTransaction)
	return len(b.chain) - 1 // returns the index of the next block to be mined
}

// Creates a new block in the Blockchain
// When blockchain is instantiated, we will need to send it to genesis block - a block with no predecessors
func first_block(proof, previousHash) {
	firstBlock = block{
		index:        1,
		timestamp:    time.Now(),
		proof:        proof,
		previousHash: 0,
	}
}

func hash(block) {
}

func main() {
	http.HandleFunc("/transactions/new", newTransaction)
	http.HandleFunc("/mine", mine)
	http.HandleFunc("/chain", fullChain)
	http.HandleFunc("/nodes/register", registerNode)
	http.HandleFunc("/nodes/resolve", consensus)
	http.ListenAndServe(":8080", nil)
}
