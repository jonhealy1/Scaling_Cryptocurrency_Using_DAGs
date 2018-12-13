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

const difficulty = 1

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

type Accounts struct {
	AccountA int
	AccountB int
	AccountC int
}

var Blockchain []Block
var BlockchainA []Block
var BlockchainB []Block
var BlockchainC []Block

type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.AccountA) + strconv.Itoa(block.AccountB) + strconv.Itoa(block.AccountC) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func generateBlock(oldBlock Block, ChangeA int, ChangeB int, ChangeC int, index string, index2 int, Account *Accounts) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.AccountA = oldBlock.AccountA + ChangeA
	newBlock.AccountB = oldBlock.AccountB + ChangeB
	newBlock.AccountC = oldBlock.AccountC + ChangeC
	Account.AccountA = Account.AccountA + ChangeA
	Account.AccountB = Account.AccountB + ChangeB
	Account.AccountC = Account.AccountC + ChangeC
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

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

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(BlockchainA) {
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

var wg sync.WaitGroup

func fork(Blocky []Block, chained int, inputData string, index2 int, ChangeA int, ChangeB int, ChangeC int, Account *Accounts) {
	defer wg.Done()

	newBlock1 := generateBlock(Blocky[len(Blocky)-1], ChangeA, ChangeB, ChangeC, inputData, index2, Account)
	if isBlockValid(newBlock1, Blocky[len(Blocky)-1]) {
		newBlockchain1 := append(Blocky, newBlock1)
		if chained == 0 {
			//mutex.Lock()
			replaceChainA(newBlockchain1)
			//	mutex.Unlock()
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

	Blockchain = append(Blockchain, genesisBlock)
	BlockchainA = append(BlockchainA, Blockchain[0])
	BlockchainB = append(BlockchainB, Blockchain[0])
	BlockchainC = append(BlockchainC, Blockchain[0])

	AccountsA := Accounts{100, 100, 100}
	AccountsB := Accounts{100, 100, 100}
	AccountsC := Accounts{100, 100, 100}

	wg.Add(3)
	inputData := "A sends B 10"
	fmt.Println("1: Account A sends Account B 10$...")
	go fork(BlockchainA, 0, inputData, 1, -10, 10, 0, &AccountsA)
	inputData2 := "B sends C 10"
	fmt.Println("2: Account B sends Account C 10$...")
	go fork(BlockchainB, 1, inputData2, 2, 0, -10, 10, &AccountsB)
	inputData3 := "3: C sends A 10"
	fmt.Println("Account C sends Account A 10$...")
	go fork(BlockchainC, 2, inputData3, 3, 10, 0, -10, &AccountsC)
	wg.Wait()
	wg.Add(3)
	inputData4 := "A sends C 5"
	fmt.Println("4: Account A sends Account C 5$...")
	go fork(BlockchainA, 0, inputData4, 4, -5, 0, 5, &AccountsA)
	inputData5 := "B sends A 5"
	fmt.Println("5: Account B sends Account A 5$...")
	go fork(BlockchainB, 1, inputData5, 5, 5, -5, 0, &AccountsB)
	inputData6 := "C sends A 20"
	fmt.Println("6: Account C sends Account A 20$...")
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
	Blockchain[len(Blockchain)-1].AccountA = BlockchainA[len(BlockchainA)-1].AccountA + BlockchainB[len(BlockchainB)-1].AccountA + BlockchainC[len(BlockchainC)-1].AccountA + Blockchain[0].AccountA - 300
	Blockchain[len(Blockchain)-1].AccountB = BlockchainA[len(BlockchainA)-1].AccountB + BlockchainB[len(BlockchainB)-1].AccountB + BlockchainC[len(BlockchainC)-1].AccountB + Blockchain[0].AccountB - 300
	Blockchain[len(Blockchain)-1].AccountC = BlockchainA[len(BlockchainA)-1].AccountC + BlockchainB[len(BlockchainB)-1].AccountC + BlockchainC[len(BlockchainC)-1].AccountC + Blockchain[0].AccountC - 300
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
