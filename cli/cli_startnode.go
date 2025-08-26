package cli

import (
	"fmt"

	"github.com/cbhcbhcbh/BlockchainInGo/chain"
)

func (cli *CLI) startNode(nodeID int) {
	fmt.Printf("Starting node %d\n", nodeID)
	chain.StartServer(nodeID)
}
