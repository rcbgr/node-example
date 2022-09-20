package main

import (
	"fmt"

	"github.com/rcbgr/node-example/pkg/node"
)

func main() {

	usdcContract := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"

	contract := "0x60e4d786628fea6478f785a6d7e704777c86a7c6"

	tokens, err := node.NodeNftApiTokensByContract(contract)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(tokens))

	fmt.Println(fmt.Sprintf("Tokens: %d", len(tokens)))

	var owners []string

	for _, token := range tokens {
		owner, err := node.NodeNftApiTokenOwnerByContractAndToken(contract, token)
		if err != nil {
			fmt.Println(err)
		}
		owners = append(owners, owner)
	}

	fmt.Println(fmt.Sprintf("Owner addresses: %d", len(owners)))

	for _, owner := range owners {

		v, err := node.NodeAdvancedApiBalanceByContract(usdcContract, owner)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("USDC balance: %f", v))

	}

}
