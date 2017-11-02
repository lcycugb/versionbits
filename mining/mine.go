package mining

import (
	"fmt"
	"time"
	"versionbits/blockchain"
)

const (
	maxNonce uint32 = 0xffffffff
)

func proofOfWork(header blockchain.BlockHeader, difficultyBits uint32) (bool, uint32) {
	targetDifficulty := blockchain.CompactToBig(difficultyBits)
	nonce := uint32(0)
	for ; nonce < maxNonce; nonce++ {
		header = header.SetNonce(nonce)
		hash := header.HashBlock()
		bigIntHash := blockchain.HashToBig(&hash)
		// fmt.Println(bigIntHash, targetDifficulty, nonce)
		if bigIntHash.Cmp(targetDifficulty) <= 0 {
			fmt.Printf("%s ", hash)
			return true, nonce
		}
	}
	return false, 0
}

// Mining yes, let's mining
func Mine() {
	fmt.Println("Mining ...")
	genesisBlock := blockchain.GetGenesisBlock()
	fmt.Println("Genesis Block: ", genesisBlock)

	prevBlock := genesisBlock

	for {
		nextBlock := prevBlock.GenerateNextBlock()
		// bits 越小难度越大
		var bits = uint32(0x20000009)
		startTime := time.Now()
		solved, nonce := proofOfWork(nextBlock, bits)
		endTime := time.Now()
		if solved {
			elapsedTime := endTime.Sub(startTime).Seconds()
			hashPower := float64(nonce) / (elapsedTime * 1000 * 1000)
			fmt.Printf("ElapsedTime: %.5f s HashPower: %.2f MH Nonce %d \n", elapsedTime, hashPower, nonce)
			prevBlock = nextBlock
		}
	}
}
