package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rcbgr/node-example/pkg/node"
)

const (
	usdcContract = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	maycContract = "0x60e4d786628fea6478f785a6d7e704777c86a7c6" // Mutant Ape Yacht Club: MAYC Token
)

func main() {

	contractTokens, err := node.NodeNftApiTokensByContract(maycContract)
	if err != nil {
		fmt.Println(err)
	}

	var owners []string
	for _, token := range contractTokens {
		if owner, err := node.NodeNftApiTokenOwnerByContractAndToken(maycContract, token); err != nil {
			fmt.Println(err)
		} else {
			owners = append(owners, owner)
		}
	}

	f, err := os.Create("balances.csv")
	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(f)
	for _, owner := range owners {
		if contractAssets, nativeAssets, err := node.NodeAdvancedApiBalanceByContract(
			usdcContract,
			owner,
		); err != nil {
			fmt.Println(err)
		} else {
			w.WriteString(fmt.Sprintf("%s,%0.18f,%f\n", owner, nativeAssets, contractAssets))
		}
	}

	w.Flush()
	f.Close()
}
