package main

import (
	"fmt"
	"log"
	"gocore/core"
	blockchain "gocore/core/blockchain"
	wallets "gocore/core/wallets"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallets.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.Close()

	UTXOSet := core.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}