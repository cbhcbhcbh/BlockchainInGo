package cli

import (
	"fmt"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) reindexUTXO() {
	bc := chain.NewBlockchain()
	UTXOSet := chain.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
