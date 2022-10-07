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
	cpContract   = "0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB" // CryptoPunks
)

func main() {

	crawlContract(cpContract, usdcContract)
	crawlContract(maycContract, usdcContract)
}

func crawlContract(nftContractAddr, usdcContractAddr string) {
	contractTokens, err := node.NodeNftApiTokensByContract(nftContractAddr)
	if err != nil {
		fmt.Println(err)
	}

	var owners []string
	for _, token := range contractTokens {
		if owner, err := node.NodeNftApiTokenOwnerByContractAndToken(nftContractAddr, token); err != nil {
			fmt.Printf("ERROR: NodeNftApiTokenOwnerByContractAndToken - token: %s - msg: %v\n", token, err)
		} else {
			owners = append(owners, owner)
		}
	}

	f, err := os.Create(fmt.Sprintf("%s-balances.csv", nftContractAddr))
	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(f)
	for _, owner := range owners {
		if contractAssets, nativeAssets, err := node.NodeAdvancedApiBalanceByContract(
			usdcContractAddr,
			owner,
		); err != nil {
			fmt.Println(err)
			fmt.Printf("ERROR: NodeAdvancedApiBalanceByContract owner: %s - msg: %v\n", owner, err)
		} else {
			w.WriteString(fmt.Sprintf("%s,%0.18f,%f\n", owner, nativeAssets, contractAssets))
		}
	}

	w.Flush()
	f.Close()

}
