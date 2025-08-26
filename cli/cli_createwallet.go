package cli

import (
	"fmt"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) createWallet() {
	wallets, _ := chain.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
