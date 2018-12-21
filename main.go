/*
The linear version of this program is based on some of the code from
this tutorial:
https://medium.com/@mycoralhealth/code-a-simple-p2p-blockchain-in-go-46662601f417
I purposefully used a very simplified blockchain implementation to explore
concurrent block creation.
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
Difficulty sets the mining difficulty to 1 which means that to
validate a block, a miner must find a valid hash for that block
starting with one zero.
*/
const difficulty = 1

/*
This ii all of the data that is stored in each Block.
Each block is storing the balance for each of Accounts
a, b, and c which is not usually done with blockchains
where actual transactions are stored instead.
*/
type Block struct {
	Index      int
	Timestamp  string
	AccountA   int
	AccountB   int
	AccountC   int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

/*
The Accounts struct is just used in tandem with the
balances stored in the Blocks for testing.
*/
type Accounts struct {
	AccountA int
	AccountB int
	AccountC int
}

/*
Each blockchain is a slice of Block structs being
linked together by the Hash and PrevHash fields.
*/
var Blockchain []Block
var BlockchainA []Block
var BlockchainB []Block
var BlockchainC []Block

/*
The calculateHash function takes all of the information in
a Block struct and calculates a hash value from it using sha256
*/
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.AccountA) + strconv.Itoa(block.AccountB) + strconv.Itoa(block.AccountC) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

/*
isHashValid approves a hash if the number of zeros at the
beginning of the hash value is equal to what is set out
in the const difficulty value
*/
func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

/*
This function serves to generate an actual block.
It appends new values to the account fields, increments
the block index, adds the timestamp and links the new
block to the old one using the last block's hash.
Furthermore it loops simulating the mining process,
continuing until the right hash is found. There is a
one second sleep added after each iteration to slow
the process down simply for viewing the output.
*/
func generateBlock(oldBlock Block, ChangeA int, ChangeB int, ChangeC int, index string, index2 int, Account *Accounts) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.AccountA = oldBlock.AccountA + ChangeA
	newBlock.AccountB = oldBlock.AccountB + ChangeB
	newBlock.AccountC = oldBlock.AccountC + ChangeC
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	Account.AccountA = Account.AccountA + ChangeA
	Account.AccountB = Account.AccountB + ChangeB
	Account.AccountC = Account.AccountC + ChangeC

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), " do more work! ", index2)
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done! ", index)
			newBlock.Hash = calculateHash(newBlock)
			break
		}
	}

	return newBlock
}

/*
isBlockValid checks to make sure that a new block is
valid and it links together with the previous block
in the chain.
*/
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

/*
The replace chain functions are used with what is
called Nakamoto consensus or the longest chain rule
to deal with forking and ensure that the longest chain
is always the valid one.
*/
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func replaceChainA(newBlocks []Block) {
	if len(newBlocks) > len(BlockchainA) {
		BlockchainA = newBlocks
	}
}

func replaceChainB(newBlocks []Block) {
	if len(newBlocks) > len(BlockchainB) {
		BlockchainB = newBlocks
	}
}

func replaceChainC(newBlocks []Block) {
	if len(newBlocks) > len(BlockchainC) {
		BlockchainC = newBlocks
	}
}

/*
waitgroup is used to ensure that all of the goroutines
representing the three parallel chains finish
before the main routine exits prematurely
*/
var wg sync.WaitGroup

/*
fork function is called via main in the goroutines
for blockchains a, b, and c and first sets off the
generateBlock function. Afterwards, if a block is
valid it is appended to the blockchain and then that
chain is checked with replaceChain to ensure that
the longest chain is being validated.
*/
func fork(Blocky []Block, chained int, inputData string, index2 int, ChangeA int, ChangeB int, ChangeC int, Account *Accounts) {
	defer wg.Done()

	newBlock1 := generateBlock(Blocky[len(Blocky)-1], ChangeA, ChangeB, ChangeC, inputData, index2, Account)
	if isBlockValid(newBlock1, Blocky[len(Blocky)-1]) {
		newBlockchain1 := append(Blocky, newBlock1)
		if chained == 0 {
			replaceChainA(newBlockchain1)
		} else if chained == 1 {
			replaceChainB(newBlockchain1)
		} else {
			replaceChainC(newBlockchain1)
		}
	}
}

func main() {
	t := time.Now()

	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 100, 100, 100, calculateHash(genesisBlock), "", difficulty, ""}

	// all of the blockchains are created with the same genesis block
	// created above with 100 in each of accounts A, B, and C.
	Blockchain = append(Blockchain, genesisBlock)
	BlockchainA = append(BlockchainA, Blockchain[0])
	BlockchainB = append(BlockchainB, Blockchain[0])
	BlockchainC = append(BlockchainC, Blockchain[0])

	AccountsA := Accounts{100, 100, 100}
	AccountsB := Accounts{100, 100, 100}
	AccountsC := Accounts{100, 100, 100}

	// first series of transactions
	wg.Add(3)
	inputData := "A sends B 10"
	fmt.Println("1: Account A sends Account B 10$...")
	time.Sleep(time.Second)
	go fork(BlockchainA, 0, inputData, 1, -10, 10, 0, &AccountsA)
	inputData2 := "B sends C 10"
	fmt.Println("2: Account B sends Account C 10$...")
	time.Sleep(time.Second)
	go fork(BlockchainB, 1, inputData2, 2, 0, -10, 10, &AccountsB)
	inputData3 := "3: C sends A 10"
	fmt.Println("Account C sends Account A 10$...")
	time.Sleep(time.Second)
	go fork(BlockchainC, 2, inputData3, 3, 10, 0, -10, &AccountsC)
	wg.Wait()

	// second deries of transactions
	wg.Add(3)
	inputData4 := "A sends C 5"
	fmt.Println("4: Account A sends Account C 5$...")
	time.Sleep(time.Second)
	go fork(BlockchainA, 0, inputData4, 4, -5, 0, 5, &AccountsA)
	inputData5 := "B sends A 5"
	fmt.Println("5: Account B sends Account A 5$...")
	time.Sleep(time.Second)
	go fork(BlockchainB, 1, inputData5, 5, 5, -5, 0, &AccountsB)
	inputData6 := "C sends A 20"
	fmt.Println("6: Account C sends Account A 20$...")
	time.Sleep(time.Second)
	go fork(BlockchainC, 2, inputData6, 6, 20, 0, -20, &AccountsC)
	wg.Wait()

	fmt.Println()
	for i := 0; i < len(BlockchainA); i++ {
		fmt.Println("Side Chain A, Block:", i, BlockchainA[i])
		fmt.Println()
		if i > 0 {
			Blockchain = append(Blockchain, BlockchainA[i])
		}
	}
	for i := 0; i < len(BlockchainB); i++ {
		fmt.Println("Side Chain B, Block:", i, BlockchainB[i])
		fmt.Println()
		if i > 0 {
			Blockchain = append(Blockchain, BlockchainB[i])
		}
	}
	for i := 0; i < len(BlockchainC); i++ {
		fmt.Println("Side Chain C, Block:", i, BlockchainC[i])
		fmt.Println()
		if i > 0 {
			Blockchain = append(Blockchain, BlockchainC[i])
		}
	}

	// account balances are settled on the main chain taking
	// information from the transactions performed on all three
	// chains.
	Blockchain[len(Blockchain)-1].AccountA = BlockchainA[len(BlockchainA)-1].AccountA + BlockchainB[len(BlockchainB)-1].AccountA + BlockchainC[len(BlockchainC)-1].AccountA + Blockchain[0].AccountA - 300
	Blockchain[len(Blockchain)-1].AccountB = BlockchainA[len(BlockchainA)-1].AccountB + BlockchainB[len(BlockchainB)-1].AccountB + BlockchainC[len(BlockchainC)-1].AccountB + Blockchain[0].AccountB - 300
	Blockchain[len(Blockchain)-1].AccountC = BlockchainA[len(BlockchainA)-1].AccountC + BlockchainB[len(BlockchainB)-1].AccountC + BlockchainC[len(BlockchainC)-1].AccountC + Blockchain[0].AccountC - 300

	// view output.
	for i := 0; i < len(Blockchain); i++ {
		fmt.Println("Main Block Chain, Block:", i, Blockchain[i])
		fmt.Println()
	}
	fmt.Println("*** Account Values on Each Sidechain ***")
	fmt.Println("* Only accounts that a sidechain is responsible for get decremented")
	fmt.Println("Sidechain A: ", AccountsA)
	fmt.Println("Sidechain B: ", AccountsB)
	fmt.Println("Sidechain C: ", AccountsC)
	fmt.Println()
	fmt.Println("*** Final Results on the Main Blockchain ***")
	fmt.Println("* These are the average results for the accounts taken from the values recorded on each sidechain")
	fmt.Println("A:", Blockchain[len(Blockchain)-1].AccountA)
	fmt.Println("B:", Blockchain[len(Blockchain)-1].AccountB)
	fmt.Println("C:", Blockchain[len(Blockchain)-1].AccountC)
}

