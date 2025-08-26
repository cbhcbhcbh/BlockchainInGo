package cli

import (
	"fmt"
	"log"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) createBlockchain(address string) {
	if !chain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := chain.CreateBlockchain(address)
	defer bc.db.Close()

	UTXOSet := chain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
