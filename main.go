package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

// block for the block chain
type block struct{
	transaction string // transaction
	nonce int // random number
	previousHash string // hash of the previous block
	currentHash string // hash of the current block
}

// the actual block chain
type blockChain struct{
	chain []*block // array of blocks pointers
	// this is basicallay dynamic silicing that contains addresses of blocks 
	// thus are linked togeher and cn traverse in list 
	// will serve as a chain of blocks 

	// the copies for genesis block and current block
	// for reference
	genesisBlock block
	currentBlock block
}

// 	function for calculating hash of a block
func  CalculateHash (stringToHash string) string {
	// calculate the hash of the given string
	sum := sha256.Sum256([]byte(stringToHash))
	hash := fmt.Sprintf("%x", sum)

	// return the hash
	return hash
}

// 	function for creating a new block and would return a block
func NewBlock(transaction string, nonce int, previousHash string) *block {
	
	// creating a new block
	// adding the information to the block
	b := new(block)
	b.transaction = transaction
	b.nonce = nonce
	b.previousHash = previousHash
	b.currentHash = ""

	// before calculating hash for the block making current hash empty
	// converting the block to string and sending it to calculate hash function
	b.currentHash = CalculateHash(fmt.Sprintf("%v", *b))

	// return the block
	return b 
}

// for creating and adding a new block to the chain
// we have a reciver for the blockChain so that we can add elements in the chain
// will return the current block
func (bc *blockChain) AddBlock(transaction string) {
	
	// creating random nounce
	nonce := rand.Intn(1000000000) 
	previousHash := ""

	// identifying the previous hash
	// if genesis block then previous hash would be empty
	if len(bc.chain) == 0 {
		previousHash = "000000000"
	}else {
		 previousHash = bc.chain[len(bc.chain)-1].currentHash
	}

	// create a new block
	b := NewBlock(transaction, nonce, previousHash)

	// add the block to the chain
	bc.chain = append(bc.chain, b)
	
	// initalizing the genesis block
	if len(bc.chain) == 1 {
		bc.genesisBlock = *b // the forst block is genesis block
	}

	// updating the current block
	bc.currentBlock = *b
}

// function to print the block chain
func (bc *blockChain) ListBlocks() {
    // printing the whole block chain
    for i, block := range bc.chain {

		
		// showing that genisis point to nothing
		if i == 0 {
			fmt.Printf("┌─────────────┐\n")
			fmt.Printf("│  000000000  │\n")
			fmt.Printf("└─────|───────┘\n")
			fmt.Printf("      | \n")
			fmt.Printf("      ^ \n")
			fmt.Printf("      | \n")
		}


		// printing the block
        fmt.Printf("┌─────|───────┐\n")
        fmt.Printf("│   Block #%d  │\n", i+1)
        fmt.Printf("├─────────────┤\n")
        fmt.Printf("│ Transaction: %v\n", block.transaction)
        fmt.Printf("│ Nonce: %v\n", block.nonce)
        fmt.Printf("│ Previous Hash: %v\n", block.previousHash)
        fmt.Printf("│ Current  Hash: %v\n", block.currentHash)

		// showing the pointer to above block
        if i != len(bc.chain)-1 {
			fmt.Printf("└─────|───────┘\n")
			fmt.Printf("      | \n")
			fmt.Printf("      ^ \n")
			fmt.Printf("      | \n")
			
		} else {
			fmt.Printf("└─────────────┘\n")
		}
	
   }
}

// 	function to change block transaction of the given block ref
func (bc *blockChain) ChangeBlock() {

	// for this we will be randomly choosing the block
	// and will change its transaction
	// for chossing randomly we are going to ranomly choose a number 
	// in length of block chain and the will alte rthe block that will
	//  be on thet iteration back from the current block
	blockToTargt := rand.Intn(len(bc.chain))
	
	// changing the transaction
	bc.chain[blockToTargt].transaction = "Changed Transaction by Trudy"
	
	 
	// as chain changed the hash of block shouls also be changed automatically
	// thus we will be calculating the hash of the block again
	bc.chain[blockToTargt].currentHash = CalculateHash(fmt.Sprintf("%v", *bc.chain[blockToTargt]))

	// thus block changed done with this function
}

// for printing the faulty blocks
// can be merged in the print given function
// but to not mixup make a new
func (bc*blockChain) faultyPrint(faultyHash string) {
	// printing the whole block chain
    for i, block := range bc.chain {
		
		// showing that genisis point to nothing
		if i == 0 {
			fmt.Printf("┌─────────────┐\n")
			fmt.Printf("│  000000000  │\n")
			fmt.Printf("└─────|───────┘\n")
			fmt.Printf("      | \n")
			fmt.Printf("      ^ \n")
			fmt.Printf("      | \n")
		}

		if(bc.chain[i].currentHash == faultyHash) {
			fmt.Printf("\033[31m")
		}
			// printing the block
			fmt.Printf("┌─────|───────┐\n")
			fmt.Printf("│   Block #%d  │\n", i+1)
			fmt.Printf("├─────────────┤\n")
			fmt.Printf("│ Transaction: %v\n", block.transaction)
			fmt.Printf("│ Nonce: %v\n", block.nonce)
			fmt.Printf("│ Previous Hash: %v\n", block.previousHash)
			fmt.Printf("│ Current  Hash: %v\n", block.currentHash)
	
			// showing the pointer to above block
			if i != len(bc.chain)-1 {
				fmt.Printf("└─────|───────┘\n")
				fmt.Printf("\033[37m")
				fmt.Printf("      | \n")
				fmt.Printf("      ^ \n")
				fmt.Printf("      | \n")
				
			} else {
				fmt.Printf("└─────────────┘\n")
				fmt.Printf("\033[37m")
			}
		
			
	
		
   }
}
// a function that will verify the chain that there is no change
// and will return true if the chain is valid
// else will return false will also display chains mean while 
func (bc * blockChain) VerifyChain() (bool) {
	// we will be iterating the whole chain and would be saving the previous hash 
	// and will be comparing it with the current hash of the next block
	// if both are same mean all good
	// that's why we had also stored the references of the genesis block and the current block
	// in relaity that would be with the many nodes and the nodes will be verifying the chain
	
	previousBlockHash :=""

	// going from the backside
	for i:=len(bc.chain)-1; i>=0; i-- {
		// the current block the last one
		if i == len(bc.chain)-1 {
			// will checkits integrity from the reference we saved 
			// as it would be with diffrent nodes
			if bc.currentBlock.currentHash != bc.chain[i].currentHash {
				fmt.Printf("\n\n\t ::::::::::::::: Chain is not valid ::::::::::::::: \n\n")
				bc.faultyPrint(bc.chain[i].currentHash)
				return false
			}
		}else if i == 0 {
			// the genesis block
			// to check the authenticity of block chain from genesis block
			// look from the reference we saved
			if bc.genesisBlock.currentHash != bc.chain[i].currentHash {
				fmt.Printf("\n\n\t ::::::::::::::: Chain is not valid ::::::::::::::: \n\n")
				bc.faultyPrint(bc.chain[i].currentHash)
				return false
			}
		}else {
		// checking the previous from successor is equal to mine or not
		if previousBlockHash != bc.chain[i].currentHash {
			fmt.Printf("\n\n\t ::::::::::::::: Chain is not valid ::::::::::::::: \n\n")
			bc.faultyPrint(bc.chain[i].currentHash)
			return false
		}

		}
		// all fine going on in chain
		previousBlockHash = bc.chain[i].previousHash

	}

	fmt.Printf("\n\n\t ::::::::::::::: Chain valid ::::::::::::::: \n\n")
	return true
}


func main(){
	
	// create a new block chain
	chain := new(blockChain)

	// hard coded transaction
	transaction := []string{"Alice sent to bob", "Trudy is here", "Another Transaction", "a small one"}

	// adding the transaction to the block chain
	for i:=0; i<len(transaction); i++ {
		chain.AddBlock(transaction[i])
	}
	
	// printing the block chain
	fmt.Printf("\n\n\t ::::::::::::::: Initial Chain ::::::::::::::: \n\n")
	chain.ListBlocks()

	fmt.Printf("\n\n\t :::::::::::::: Changing Chain ::::::::::::::: \n\n")

	// changing the block
	chain.ChangeBlock() 

	// printing the block chain
	fmt.Printf("\n\n\t ::::::::::::::: Verifying Chain ::::::::::::::: \n\n")

	// verifying the chain
	if ( chain.VerifyChain() ){
		chain.ListBlocks()
	}

	fmt.Println()
	fmt.Println()
	
}