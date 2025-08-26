package cli

import (
	"fmt"
	"log"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) getBalance(address string) {
	if !chain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := chain.NewBlockchain()
	UTXOSet := chain.UTXOSet{bc}
	defer bc.db.Close()

	balance := 0
	pubKeyHash := chain.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
