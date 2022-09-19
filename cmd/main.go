package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"

	nodec "github.com/rcbgr/node-example/pkg/node"
	"github.com/rcbgr/node-example/pkg/node/model"
)

func main() {

	config := &model.NodeConfig{
		Username: os.Getenv("NODE_USERNAME"),
		Password: os.Getenv("NODE_PASSWORD"),
		Endpoint: os.Getenv("NODE_ENDPOINT"),
	}

	req := &model.NodeAdvancedApiRequest{
		Config: config,
		Body: &model.JsonRpcBody{
			Id:             uuid.New().String(),
			JsonRpcVersion: "2.0",
			Method:         "coinbaseCloud_getSingleBalance",
			Params: &model.NodeAdvancedApiGetSingleBalanceRequest{
				Address:    "0x0a59649758aa4d66e25f08dd01271e891fe52199", // Maker: PSM-USDC-A
				Contract:   "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC Contract
				Blockchain: "Ethereum",
				Network:    "Mainnet",
			},
		},
	}

	res, err := nodec.NodeAdvancedApiCall(req)
	if err != nil {
		log.Fatalf("Failed to call coinbaseCloud_getSingleBalance: %v", err)
	}

	fmt.Println(string(res.Body))

}
