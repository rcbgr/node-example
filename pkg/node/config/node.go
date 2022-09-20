package config

import (
	"os"
	"strings"

	"github.com/rcbgr/node-example/pkg/node/model"
)

var NodeConfig *model.NodeConfig

func init() {

	NodeConfig = &model.NodeConfig{
		Username: os.Getenv("NODE_USERNAME"),
		Password: os.Getenv("NODE_PASSWORD"),
		Endpoint: os.Getenv("NODE_ENDPOINT"),
	}

	if strings.Contains(strings.ToLower(NodeConfig.Endpoint), "ethereum") {
		NodeConfig.Blockchain = "Ethereum"
	}

	if strings.Contains(strings.ToLower(NodeConfig.Endpoint), "mainnet") {
		NodeConfig.Network = "Mainnet"
	}

	if strings.Contains(strings.ToLower(NodeConfig.Endpoint), "ethereum") &&
		strings.Contains(strings.ToLower(NodeConfig.Endpoint), "mainnet") {
		NodeConfig.NftNetworkName = "ethereum-mainnet"
	}
}
