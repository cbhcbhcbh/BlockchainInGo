package cli

import (
	"fmt"
	"log"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) send(from, to string, amount int) {
	if !chain.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !chain.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := chain.NewBlockchain()
	UTXOSet := chain.UTXOSet{bc}
	defer bc.db.Close()

	tx := chain.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := chain.NewCoinbaseTX(from, "")
	txs := []*chain.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")
}
